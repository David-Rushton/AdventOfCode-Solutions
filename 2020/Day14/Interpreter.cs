using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Interpreter
    {
        readonly Tokeniser _tokeniser;
        readonly Converter _converter;
        readonly Dictionary<int, string> _memory = new ();
        readonly string _defaultValue = "".PadLeft(36, '0');


        public Interpreter(Tokeniser tokeniser, Converter converter) =>
            (_tokeniser, _converter) = (tokeniser, converter)
        ;


        public void Invoke()
        {
            foreach(var token in _tokeniser.GetTokens())
            {
                Verbose.WriteLine($"  Processing token: {token}");

                InitialiseMemoryIfNotExists(token.Address);
                UpdateMemory(token.Address, token.Mask, token.Value);
            }


            Console.WriteLine($"\nResult: {GetResult()}\n");


            void InitialiseMemoryIfNotExists(int address)
            {
                if( ! _memory.ContainsKey(address) )
                    _memory.Add(address, _defaultValue);
            }

            void UpdateMemory(int address, string mask, long value) =>
                _memory[address] = _converter.ApplyMask(_memory[address], mask, value)
            ;

            long GetResult() =>
                _memory
                    .Where(kvp => kvp.Value != _defaultValue)
                    .Select(kvp => _converter.ToLong(kvp.Value))
                    .Sum()
            ;

        }
    }
}
