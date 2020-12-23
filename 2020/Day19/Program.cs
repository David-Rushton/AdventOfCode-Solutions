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

            Console.WriteLine(components.Parser.GetRuleZero(components.Lexer.GetRules()));
            Environment.Exit(0);


            bool UseTestInput() => args.Contains("--test");

            string GetInputPath() =>
                Path.Join(Directory.GetCurrentDirectory(), UseTestInput() ? "Test-Input.txt" : "Input.txt")
            ;

            (Lexer Lexer, Parser Parser) Bootstrap() =>
                (
                    new Lexer(GetInputPath()),
                    new Parser()
                )
            ;
        }
    }
}
