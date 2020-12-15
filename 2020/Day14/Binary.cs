using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Binary
    {
        public Binary(int addressSpace) => (AddressSpace) = (addressSpace);


        public int AddressSpace { get; init; }


        public string To(long Value) =>
            new string(System.Convert.ToString(Value, 2).PadLeft(AddressSpace, '0').ToCharArray())
        ;

        public long From(string Value) => From(Value.ToCharArray());
        public long From(char[] Value)
        {
            long result = 0;
            long positionValue = 1;

            for(var i = Value.Length -1; i >= 0; i--)
            {
                if(Value[i] == '1')
                    result += positionValue;

                positionValue = positionValue * 2;
            }

            return result;
        }
    }
}
