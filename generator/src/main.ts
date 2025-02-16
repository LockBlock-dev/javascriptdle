import * as cheerio from "cheerio";
import fs from "node:fs/promises";

const UNWANTED_FILTER = ["get", "set", "iterator", "%"];
const MUST_CONTAIN_FILTER = ["(", ")"];
const PARAM_REGEX = /,?_\w+_/g;
const OPTIONAL_PARAM_REGEX = /\[,?_\w+_/g;

const main = async () => {
    // https://raw.githubusercontent.com/tc39/ecma262/refs/heads/main/spec.html
    const specFile = await fs.readFile("./spec.html");
    const $ = cheerio.load(specFile);
    const data = (
        $("emu-clause")
            .not("[type]")
            .map((_, el) => $(el).children("h1").first().text())
            .get() as Array<string>
    )
        .filter(
            (el) =>
                !UNWANTED_FILTER.some((filter) => el.includes(filter)) &&
                MUST_CONTAIN_FILTER.every((filter) => el.includes(filter)) &&
                el.split("(")[0].includes("."),
        )
        .map((el) => {
            const [first, ...paramsArray] = el.split(" ");
            const methodMembers = first.split(".");
            const params = paramsArray.join("");

            const totalParams = params.match(PARAM_REGEX)?.length ?? 0;
            const optionalParams =
                params.match(OPTIONAL_PARAM_REGEX)?.length ?? 0;

            return {
                object: methodMembers.at(0),
                method: methodMembers.at(-1),
                isPrototype: first.includes("prototype"),
                mandatoryParams: totalParams - optionalParams,
                optionalParams,
                totalParams,
            };
        });

    await fs.writeFile("./data.json", JSON.stringify(data, null, 2), {
        encoding: "utf8",
    });
};

main();
