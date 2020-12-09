using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var useTestInput = args.Contains("--use-test-input");
            var verbose = args.Contains("--verbose");
            var parser = new Parser();
            var instructions = parser.Parse(useTestInput);
            var interpreter = new Interpreter();
            var corruptionCorrector = new CorruptionCorrector();
            var fixCorruption = args.Contains("--fix-corruption");
            var maxAttempts = fixCorruption ? 100 : 1;
            var attempts = 0;
            InterpreterResult result;


            while(true)
            {
                attempts++;
                result = interpreter.Invoke(instructions, verbose);
                Console.WriteLine($"\nInterpreter complete\nAttempt: {attempts}\nExit reason: {result.ExitCode}\nResult: {result.Accumulator}");

                if(maxAttempts == 1 || attempts >= maxAttempts || result.ExitCode == InterpreterExitCodes.Success)
                    break;

                instructions = corruptionCorrector.AttemptCorrection(result.Instructions, verbose);
            }
        }
    }
}
