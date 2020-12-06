using System;
using System.Diagnostics;
using System.Text.RegularExpressions;


namespace AoC
{
    public record PasswSord(
        char RequiredChar,
        int MinRequiredInstances,
        int MaxRequiredInstances,
        string Pwd
    );


    public class PasswordParser
    {
        // Format:  10-11 r: rrrxrrrrwhrrrr
        //         |-------||--------------|
        //         |policy ||password      |
        // Where:
        //  10 is min instances of r
        //  11 is max instances of r
        //  rrrxrrrrwhrrrr is password
        const string ExpectedInputPattern = @"\d+-\d+\s{1}[a-z]{1}:\s{1}[a-z]+";

        const string CannotParsePasswordMessage = "Cannot parse password unexpected format: {0}";


        public Password Parser(string input)
        {
            Debug.Assert(
                new Regex(ExpectedInputPattern).IsMatch(input),
                string.Format(CannotParsePasswordMessage, input)
            );

            var inputElements = input.Split(':');
            var policyElements = inputElements[0].Split(' ');
            var policyChecks = policyElements[0].Split('-');
            var policyCheckOne = int.Parse(policyChecks[0]);
            var policyCheckTwo = int.Parse(policyChecks[1]);
            var policyRequiredChar = policyElements[1].ToCharArray()[0];
            var pwd = inputElements[1].Trim();


            return new Password(policyRequiredChar, policyCheckOne, policyCheckTwo, pwd);
        }
    }
}
