using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    class Program
    {
        const int _addressSpace = 36;


        static void Main(string[] args)
        {
            var useV2 = args.Contains("--v2");
            var useTestInput = args.Contains("--test");
            var showVerboseOutput = args.Contains("--verbose");
            var interpreter = Bootstrap(useV2, useTestInput, showVerboseOutput);

            interpreter.Invoke();
        }


        private static Interpreter Bootstrap(bool useV2, bool useTestInput, bool showVerboseMode)
        {
            Verbose.ShowVerboseOutput = showVerboseMode;

            var inputPath = Path.Join(Directory.GetCurrentDirectory(), useTestInput ? "Test-Input.txt" : "Input.txt");
            var binary = new Binary(_addressSpace);
            IMemoryAddressDecoder memoryAddressDecoder =
                useV2
                ? new MemoryAddressDecoderV2(binary)
                : new MemoryAddressDecoderV1()
            ;
            var tokeniser = new Tokeniser(inputPath);
            var converter = new Converter(binary);
            var interpreter = new Interpreter(useV2, tokeniser, converter, memoryAddressDecoder);

            return interpreter;
        }
    }
}
