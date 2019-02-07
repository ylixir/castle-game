#!/usr/bin/env bash

nixProfile="$HOME/.nix-profile/etc/profile.d/nix.sh"

source "$nixProfile"

#####################
## dev environment ##
#####################

function setup-nix {
    if [ -x "$(command -v nix)" ]; then
        echo "Nix is already installed."
        echo "You may update it with the following commands:"
        echo "  $ nix-channel --update"
        echo "  $ nix-env -u"
    else
        function gpg_fake {
            echo "Faking GPG verification"
            true
        }

        if [ -x "$(command -v gpg2)" ]; then
            gpg=gpg2
        else
            echo "WARNING!! GPG not installed, cannot verify authenticity of nix!"
            echo "Press ^C to cancel, or any other key to continue."
            read -n 1 -s
            gpg=gpg_fake
        fi

        mkdir -p nixtmp
        curl -o nixtmp/install-nix https://nixos.org/nix/install \
        && curl -o nixtmp/install-nix.sig https://nixos.org/nix/install.sig \
        && $gpg --recv-keys B541D55301270E0BCF15CA5D8170B4726D7198DE \
        && $gpg --verify nixtmp/install-nix.sig \
        && sh nixtmp/install-nix
        rm -rf nixtmp
    fi
}

#execute a shell with debugging disabled
function run {
    if [ $# -eq 0 ]
    then
        nix run
    else
        nix run -c "${@}"
    fi
}

########################
## project management ##
########################


function init {
    setup-nix
    source "$nixProfile"
    run npm install
}

function version {
    grep version package.json | sed "s/^[^0-9]*\([0-9.]*\).*$/v\\1/g"
}
function tagged {
    git tag --merged master | grep $(version) > /dev/null
}

function tag {
    tagged || git tag $(version)
    git push origin $(version)
}


#####################
## quality control ##
#####################

function npm {
    run npm "${@}"
}

function format-verify {
    npm run format -- --list-different
}

function format {
    npm run format -- --write
}

function build {
    npm run build
}

function clean {
    npm run clean
}

function check {
    build \
    && format
}


$1 ${@:2}