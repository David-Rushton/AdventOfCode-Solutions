using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var components = Bootstrap(args);
            components.analyser.Analyse(components.foods, components.ingredients, components.allergens);
        }


        private static (Analyser analyser, List<Food> foods, Dictionary<string, int> ingredients, Dictionary<string, int> allergens) Bootstrap(
            string[] args
        )
        {
            var tokens = new Lexer().GetTokens(args.Contains("--test"));

            return
            (
                new Analyser(),
                tokens.foods,
                tokens.ingredients,
                tokens.allergens
            );
        }
    }
}
