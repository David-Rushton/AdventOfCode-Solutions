using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class MemoryAddressDecoderV2: IMemoryAddressDecoder
    {
        readonly char[] _overrideCharacters = new [] { '1', 'X' };
        readonly char[] _resolveCharacters = new [] { '0', '1' };
        readonly Binary _binary;


        public MemoryAddressDecoderV2(Binary binary) => (_binary) = (binary);


        public IEnumerable<long> Decode(int address, string mask) =>
            ResolveFloatingBits(ApplyMask(_binary.To(address), mask))
        ;


        private string ApplyMask(string value, string mask)
        {
            var result = string.Empty;

            for(var i = 0; i < value.Length; i++)
            {
                result +=
                        _overrideCharacters.Contains(mask[i])
                    ? mask[i]
                    : value[i]
                ;
            }

            return result;
        }

        private IEnumerable<long> ResolveFloatingBits(string value)
        {
            var firstIndexOfX = value.IndexOf('X');

            if(firstIndexOfX >= 0)
            {
                foreach(var replaceChar in _resolveCharacters)
                    foreach(var item in ResolveFloatingBits(replaceFirstX(replaceChar)))
                        yield return item;
            }
            else
            {
                Verbose.WriteLine($"  Resolved floating point: {value}");
                yield return _binary.From(value);
            }


            string replaceFirstX(char newValue) =>
                value.Substring(0, firstIndexOfX) + newValue + value.Substring(firstIndexOfX + 1);
            ;
        }
    }
}
