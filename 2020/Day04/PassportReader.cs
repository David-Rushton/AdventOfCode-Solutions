using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public class PassportReader
    {
        readonly string _inputFilePath;


        public PassportReader(string inputFilePath) => (_inputFilePath) = (inputFilePath);


        public IEnumerable<Passport> GetPassports()
        {
            var nextPassport = new Passport();

            foreach(var line in File.ReadLines(_inputFilePath))
            {
                if(string.IsNullOrEmpty(line))
                {
                    yield return nextPassport;
                    nextPassport = new Passport();
                    continue;
                }

                var fields = line.Split(' ');
                foreach(var field in fields)
                {
                    var kv = field.Split(':');
                    nextPassport.Fields.Add(kv[0], kv[1]);
                }
            }

            yield return nextPassport;
        }
    }
}
