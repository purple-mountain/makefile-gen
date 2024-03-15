const fs = require("fs")
const program_name = process.argv.slice(2)[0]

if (!program_name) {
  throw new Error("Invalid program name")
}

const makefileContent = `${program_name}: ${program_name}.c
\tg++ -o ./bin/${program_name} ${program_name}.c`

fs.writeFile("makefile", makefileContent, (err) => {
  if (err) throw err;
})
