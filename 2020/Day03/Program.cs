using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public static class Program
    {
        const string UnexpectedArgExceptionMessage = "Expected arg --star-one or --star-two";


        public static void Main(string[] args)
        {
            if(args.Length != 1)
                throw new Exception(UnexpectedArgExceptionMessage);

            var (star, input) = Bootstrap(args[0]);
            star.Invoke(input);
        }


        private static (IStar star, List<string> input) Bootstrap(string starArg) =>
            (GetStar(starArg), GetInput())
        ;

        private static IStar GetStar(string starArg) =>
            starArg switch
            {
                "--star-one" => new StarOne(),
                "--star-two" => new StarTwo(),
                _            => throw new Exception(UnexpectedArgExceptionMessage),
            };

        private static List<string> GetInput()
        {
            var path = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");
            return File.ReadLines(path).ToList<string>();
        }
    }
}
