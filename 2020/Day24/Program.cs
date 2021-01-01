using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var (inputPath, lexer, tiler) = Bootstrap(args.Contains("--test"));

            tiler.Lay(lexer.GetTokens(inputPath));
        }


        private static (string inputPath, Lexer lexer, Tiler tiler) Bootstrap(bool useTestInput) =>
            (
                Path.Join(Directory.GetCurrentDirectory(), useTestInput ? "TestInput.txt" : "Input.txt"),
                new Lexer(),
                new Tiler()
            )
        ;
    }
}
