using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    public class CorruptionCorrector
    {
        readonly List<int> _correctedInstructions = new();

        int _mutateIndex = -1;


        public List<Instruction> AttemptCorrection(List<Instruction> instructions, bool verbose)
        {
            ResetLastMutation();


            _mutateIndex = instructions
                .Where(i => (i.Visited == true && (i.Operation is Operations.Jump or Operations.NoOperation)))
                .Max
                (
                    i =>
                    {
                        if( ! _correctedInstructions.Contains(i.Index) )
                            return i.Index;

                        return -1;
                    }
                );


            _correctedInstructions.Add(_mutateIndex);
            FlipInstruction();
            IfVerboseReportMutation();


            return ResetVisited();


            // Only change can be made at a time.
            // Removes last change ahead of next.
            void ResetLastMutation()
            {
                if(_mutateIndex >= 0)
                    FlipInstruction();
            }

            // Flips a jump to no-op or a no-op to a jump.
            void FlipInstruction() =>
                instructions[_mutateIndex] = instructions[_mutateIndex] with
                {
                    Operation = FlipOperation()
                }
            ;

            Operations FlipOperation() =>
                instructions[_mutateIndex].Operation is Operations.Jump ? Operations.NoOperation : Operations.Jump
            ;

            // In verbose mode we output the mutation to the screen.
            void IfVerboseReportMutation()
            {
                if(verbose)
                    Console.WriteLine($"\nInstruction mutated: {instructions[_mutateIndex]}");
            }

            List<Instruction> ResetVisited() =>
                instructions.Select(i => i with { Visited = false }).ToList()
            ;
        }
    }
}
