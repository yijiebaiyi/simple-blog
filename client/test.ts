import { remark } from "remark";

main();

async function main() {
  const file = await remark()
    .data("settings", { bullet: "*", setext: true, listItemIndent: "one" })
    .process("# Moons of Neptune\n\n- Naiad\n- Thalassa\n- Despine\n- …");

  console.log(String(file));
}
