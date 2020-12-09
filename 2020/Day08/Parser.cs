using System;
using System.Collections.Generic;
using System.IO;


namespace AoC
{

    public enum Operations
    {
        Accumulator,
        Jump,
        NoOperation
    }


    public record Instruction(
        int Index,
        Operations Operation,
        int Argument,
        bool Visited
    );


    public class Parser
    {
        readonly string _testInput = Path.Join(Directory.GetCurrentDirectory(), "Test-Input.txt");

        readonly string _inputPath = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");


        public List<Instruction> Parse(bool useTestInput)
        {
            var index = 0;
            var instructions = new List<Instruction>();
            var inputPath = useTestInput ? _testInput : _inputPath;

            foreach(var line in File.ReadLines(inputPath))
            {
                var elements = line.Split(' ');
                instructions.Add
                (
                    new Instruction
                    (
                        index++,
                        elements[0] switch
                        {
                            "acc" => Operations.Accumulator,
                            "jmp" => Operations.Jump,
                            "nop" => Operations.NoOperation,
                            _ => throw new Exception($"Unsupported operation: {elements[0]}")
                        },
                        int.Parse(elements[1]),
                        false
                    )
                );
            }


            return instructions;
        }
    }
}
