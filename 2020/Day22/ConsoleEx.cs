using System;


namespace AoC
{
    public class ConsoleEx
    {
        public static void WriteLine(string message, int game = 1)
        {
            if(game == 1)
                Console.WriteLine(message);
        }
    }
}
