open Reprocessing;

type stateT = {
  shrek: imageT
}
let setup = (env) => {
  Env.size(~width=600, ~height=600, env);
  let shrek = Draw.loadImage(
    ~filename="assets/ogre.jpg"
    , env
    );
  {
    shrek
  }
}

let draw = (state, env) => {
  Draw.background(Utils.color(~r=100, ~g=200, ~b=255, ~a=255), env);
  Draw.fill(Utils.color(~r=150, ~g=255, ~b=75, ~a=255), env);
  Draw.rect(~pos=(1, 450), ~width=600, ~height=150, env);
  Draw.image(state, ~pos=(0,100),~width=135, ~height=500,env);

  state
}

run(~setup, ~draw, ());