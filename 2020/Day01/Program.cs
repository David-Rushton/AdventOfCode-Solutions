using System;


namespace Day01
{
    public static class Program
    {
        public static void Main(string[] args)
        {
            if(args[0] == "--star-one")
                new StarOne().Invoke();

            if(args[0] == "--star-two")
                new StarTwo().Invoke();
        }
    }
}
