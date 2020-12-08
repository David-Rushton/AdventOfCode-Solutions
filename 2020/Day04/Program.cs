using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;
using AoC.PassportRules;


namespace AoC
{
    public static class Program
    {
        const string UnexpectedArgExceptionMessage = "Expected arg --star-one or --star-two";

        static string _inputPath => Path.Join(Directory.GetCurrentDirectory(), "Input.txt");


        public static void Main(string[] args)
        {
            if(args.Length != 1)
                throw new Exception(UnexpectedArgExceptionMessage);

            var (star, reader) = Bootstrap(args[0]);
            star.Invoke(reader);
        }


        private static (IStar star, PassportReader reader) Bootstrap(string starArg) =>
            (
                GetStar(starArg),
                new PassportReader
                (
                    _inputPath,
                    new IPassportRule[]
                    {
                        new BirthYearPassportRule(),
                        new ExporationYearPassportRule(),
                        new EyeColourPassportRule(),
                        new HairColourPassportRule(),
                        new HeightPassportRule(),
                        new IssueYearPassportRule(),
                        new PassportIdPassportRule()
                    }
                )
            )
        ;

        private static IStar GetStar(string starArg) =>
            starArg switch
            {
                "--star-one" => new StarOne(),
                "--star-two" => new StarTwo(),
                _            => throw new Exception(UnexpectedArgExceptionMessage),
            };
    }
}
