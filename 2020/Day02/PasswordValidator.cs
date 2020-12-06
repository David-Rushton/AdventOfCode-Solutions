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

        public bool IsTobogganValidPassword(string input)
        {
            var pwd = _passwordParser.Parser(input);
            var pwdCharArray = pwd.Pwd.ToCharArray();
            var requireCharCount = 0;


            foreach(var checkPosition in new [] {pwd.CheckOne, pwd.CheckTwo})
            {
                // WARNING: Position to check is one based.
                //          pwdChar is zero based.
                if(checkPosition <= pwdCharArray.Length && pwdCharArray[checkPosition - 1] == pwd.RequiredChar)
                    requireCharCount++;
            }


            return (requireCharCount == 1);
        }
    }
}
