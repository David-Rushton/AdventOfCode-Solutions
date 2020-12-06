using System;


namespace Day01
{
    public static class Program
    {
        const string UnexpectedArgExceptionMessage = "Expected arg --star-one or --star-two";


        public static void Main(string[] args)
        {
            if(args.Length != 1)
                throw new Exception(UnexpectedArgExceptionMessage);

            GetStar(args[0]).Invoke();
        }


        public static IStar GetStar(string starArg) =>
            starArg switch
            {
                "--star-one" => new StarOne(),
                "--star-two" => new StarTwo(),
                _            => throw new Exception(UnexpectedArgExceptionMessage),
            };
    }
}
