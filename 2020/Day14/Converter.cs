using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Converter
    {
        readonly char[] _overrideCharacters = new [] {'0', '1'};


        public string ApplyMask(string from, string mask, long value)
        {
            var len = from.Length;
            var to = System.Convert.ToString(value, 2).PadLeft(len, '0').ToCharArray();

            for(var i = 0; i < len; i++)
            {
                if(_overrideCharacters.Contains(mask[i]))
                {
                    to[i] = mask[i];
                }
            }


            return new string(to);
        }

        public long ToLong(string from)
        {
            long result = 0;
            long value = 1;

            for(var i = from.Length - 1; i >= 0; i--)
            {
                if(from[i] == '1')
                   result += value;

                value = value * 2;
            }


            Verbose.WriteLine($"  Binary {from} {result}");
            return result;
        }
    }
}
