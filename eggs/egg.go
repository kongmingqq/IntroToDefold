components {
  id: "egg"
  component: "/eggs/egg.script"
}
components {
  id: "pickup"
  component: "/eggs/pickup.particlefx"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"egg\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/eggs/eggs.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.1\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "mask: \"sensor\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: -12.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"egg_btm\"\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 15.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 1\n"
  "    count: 1\n"
  "    id: \"egg_top\"\n"
  "  }\n"
  "  data: 32.216522\n"
  "  data: 26.248844\n"
  "}\n"
  ""
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/impactWood_heavy_003.ogg\"\n"
  ""
}
