using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Interpreter
    {
        readonly bool _overwrite;
        readonly Tokeniser _tokeniser;
        readonly Converter _converter;
        readonly Dictionary<long, string> _memory = new ();
        readonly IMemoryAddressDecoder _memoryAddressDecoder;
        readonly string _defaultValue = "".PadLeft(36, '0');


        public Interpreter(bool overwrite, Tokeniser tokeniser, Converter converter, IMemoryAddressDecoder memoryAddressDecoder) =>
            (_overwrite, _tokeniser, _converter, _memoryAddressDecoder) = (overwrite, tokeniser, converter, memoryAddressDecoder)
        ;


        public void Invoke()
        {
            foreach(var token in _tokeniser.GetTokens())
            {
                Verbose.WriteLine($"  Processing token: {token}");

                foreach(var address in _memoryAddressDecoder.Decode(token.Address, token.Mask))
                {
                    InitialiseMemoryIfNotExists(address);
                    UpdateMemory(address, token.Mask, token.Value);
                }
            }


            Console.WriteLine($"\nResult: {GetResult()}\n");


            void InitialiseMemoryIfNotExists(long address)
            {
                if( ! _memory.ContainsKey(address) )
                    _memory.Add(address, _defaultValue);
            }

            void UpdateMemory(long address, string mask, long value)
            {
                if(_overwrite)
                {
                    _memory[address] = System.Convert.ToString(value, 2).PadLeft(36, '0');
                }
                else
                {
                    _memory[address] = _converter.ApplyMask(_memory[address], mask, value);
                }
            }

            long GetResult() =>
                _memory
                    .Where(kvp => kvp.Value != _defaultValue)
                    .Select(kvp => _converter.ToLong(kvp.Value))
                    .Sum()
            ;

        }
    }
}
