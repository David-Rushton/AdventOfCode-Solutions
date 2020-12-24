using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;


namespace AoC
{
    public class Interpreter
    {
        public int CountOfImagesThatMatchRule(string rule, IEnumerable<string> images) =>
            (

                from image in images
                let ruleRegex = new Regex(rule)
                where ruleRegex.IsMatch(image)
                select image
            ).Count()
        ;
    }
}
