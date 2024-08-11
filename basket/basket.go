components {
  id: "basket"
  component: "/basket/basket.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"basket\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/player_sprites.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "edges"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -96.0\n"
  "      y: -54.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: 0.16943027\n"
  "      w: 0.9855422\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"basket_left\"\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 95.0\n"
  "      y: -55.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: -0.13847591\n"
  "      w: 0.9903658\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "    id: \"basket_right\"\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: -126.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "    id: \"basket_bottom\"\n"
  "  }\n"
  "  data: 6.4430227\n"
  "  data: 69.2187\n"
  "  data: 10.0\n"
  "  data: 6.4430227\n"
  "  data: 69.2187\n"
  "  data: 10.0\n"
  "  data: 80.344826\n"
  "  data: 3.181361\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sensor"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"sensor\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -10.0\n"
  "      y: -93.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"sensorshape\"\n"
  "  }\n"
  "  data: 27.174974\n"
  "}\n"
  ""
}
