using System;
using System.Collections.Generic;


namespace AoC
{
    public class Interpreter
    {
        public int Invoke(List<Instruction> instructions, bool verbose)
        {
            var accumulator = 0;
            var index = 0;
            var instructionsProcessed = 0;

            while(true)
            {
                var instruction = instructions[index];

                // We've entered an infinite loop!
                if(instruction.Visited)
                    break;

                // Flag as visited before we change index
                instructions[index] = instruction with {Visited = true};

                switch (instruction.Operation)
                {
                    case Operations.Accumulator:
                        accumulator += instruction.Argument;
                        index++;
                        break;

                    case Operations.Jump:
                        index += instruction.Argument;
                        break;

                    case Operations.NoOperation:
                        index++;
                        break;

                    default:
                        throw new Exception($"Unsupported instruction: {instruction}");
                }

                instructionsProcessed++;

                if(verbose)
                    Console.WriteLine($"{instructionsProcessed.ToString("##")}: {instruction}");
            }

            return accumulator;
        }
    }
}
