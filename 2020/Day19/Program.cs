using System;
using System.Linq;
using System.IO;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var components = Bootstrap();
            var rulesInput = components.Lexer.GetRules();
            var rule = components.Parser.GetRuleZero(rulesInput);
            var result = components.Interpreter.CountOfImagesThatMatchRule(rule, components.Lexer.GetImages());


            Console.WriteLine($"\nMatching images: {result}");
            Environment.Exit(0);


            bool UseTestInput() => args.Contains("--test");

            string GetInputPath() =>
                Path.Join(Directory.GetCurrentDirectory(), UseTestInput() ? "Test-Input.txt" : "Input.txt")
            ;

            (Lexer Lexer, Parser Parser, Interpreter Interpreter) Bootstrap() =>
                (
                    new Lexer(GetInputPath()),
                    new Parser(),
                    new Interpreter()
                )
            ;
        }
    }
}
