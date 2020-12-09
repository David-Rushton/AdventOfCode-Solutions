using System;
using System.Collections.Generic;


namespace AoC
{
    public enum InterpreterExitCodes
    {
        Success = 0,
        Overflow = 1,
        InfiniteLoop = 2
    };


    public class InterpreterResult
    {
        public int Accumulator { get; set; }
        public InterpreterExitCodes ExitCode { get; set; }

        #nullable disable
        public List<Instruction> Instructions { get; set; }
        #nullable enable
    }


    public class Interpreter
    {
        public InterpreterResult Invoke(
            List<Instruction> instructions,
            bool verbose
        )
        {
            var accumulator = 0;
            var index = 0;


            while( ! programComplete() )
            {
                var instruction = instructions[index];


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

                if(verbose)
                    Console.WriteLine(instruction);
            }


            return new InterpreterResult
            {
                Accumulator =  accumulator,
                ExitCode = getExitCode(),
                Instructions = instructions
            };


            bool programComplete() => index == instructions.Count || instructions[index].Visited;

            InterpreterExitCodes getExitCode() =>
                index switch
                {
                    int idx when idx == instructions.Count  => InterpreterExitCodes.Success,
                    int idx when idx > instructions.Count   => InterpreterExitCodes.Overflow,
                    _                                       => InterpreterExitCodes.InfiniteLoop
                }
                // index == instructions.Count ? InterpreterExitCodes.Success : InterpreterExitCodes.InfiniteLoop
            ;
        }
    }
}
