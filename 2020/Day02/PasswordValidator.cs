using System;
using System.Linq;


namespace AoC
{
    public class PasswordValidator
    {
        readonly PasswordParser _passwordParser;


        public PasswordValidator(PasswordParser passwordParser) => (_passwordParser) = (passwordParser);


        public bool IsSledValidPassword(string input)
        {
            var pwd = _passwordParser.Parser(input);
            var requiredCharCount = pwd.Pwd.Count(c => c == pwd.RequiredChar);
            var isValid = (requiredCharCount >= pwd.CheckOne && requiredCharCount <= pwd.CheckTwo);

            return isValid;
        }
    }
}
