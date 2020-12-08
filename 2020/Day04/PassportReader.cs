using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using AoC.PassportRules;


namespace AoC
{
    public record PassportReaderStatistics(
        int countOfProcessedPassports,
        int countOfPassportsWithAllRequiredFields,
        int countOfPassportsThatPassAllRules
    );


    public class PassportReader
    {
        readonly string _inputFilePath;

        readonly Dictionary<string, IPassportRule> _passportRules = new Dictionary<string, IPassportRule>();

        readonly int _requiredFieldCount;


        public PassportReader(string inputFilePath, IPassportRule[] passportRules)
        {
            _inputFilePath = inputFilePath;

            _requiredFieldCount = passportRules.Length;
            foreach(var rule in passportRules)
            {
                _passportRules.Add(rule.FieldName, rule);
            }
        }


        public PassportReaderStatistics ProcessPassports()
        {
            int countOfProcessedPassports = 0;
            int countOfPassportsWithAllRequiredFields = 0;
            int countOfPassportsThatPassAllRules = 0;

            foreach(var passport in ValidatePassports())
            {
                countOfProcessedPassports++;

                if(passport.hasAllRequiredFields)
                    countOfPassportsWithAllRequiredFields++;

                if(passport.allRequiredFieldsValidated)
                    countOfPassportsThatPassAllRules++;
            }

            return new PassportReaderStatistics(
                countOfProcessedPassports,
                countOfPassportsWithAllRequiredFields,
                countOfPassportsThatPassAllRules
            );
        }


        private IEnumerable<(bool hasAllRequiredFields, bool allRequiredFieldsValidated)> ValidatePassports()
        {
            foreach(var passport in GetPassports())
            {
                int requiredFieldsCount = 0;
                int validatedRequiredFieldCount = 0;

                foreach(var field in GetPassportFields(passport))
                {
                    if(field.isRequiredField)
                        requiredFieldsCount++;

                    if(field.containsValidValue)
                        validatedRequiredFieldCount++;
                }

                yield return (requiredFieldsCount == _requiredFieldCount, validatedRequiredFieldCount == _requiredFieldCount);
            }
        }

        private IEnumerable<(string fieldName, string fieldValue, bool isRequiredField, bool containsValidValue)> GetPassportFields(string passport) =>
            passport.Split(' ')
            .Where(f => f.Count(c => c ==':') == 1)
            .Select
            (
                f =>
                {
                    var elements = f.Split(':');
                    var fieldName = elements[0].Trim();
                    var fieldValue = elements[1].Trim();
                    var testResult = TestPassportField(fieldName, fieldValue);

                    return (fieldName, fieldValue, testResult.isRequiredField, testResult.containsValidValue);
                }
            )
        ;

        private (bool isRequiredField, bool containsValidValue) TestPassportField(string fieldName, string fieldValue)
        {
            var isRequiredField = _passportRules.ContainsKey(fieldName);
            var containsValidValue = false;

            // all required fields have an associated rule.
            if(isRequiredField)
                containsValidValue = _passportRules[fieldName].IsValidValue(fieldValue);

            return (isRequiredField, containsValidValue);
        }

        // Forward only passport reader.
        private IEnumerable<string> GetPassports()
        {
            var passport = string.Empty;
            var newlineAndTabRegEx = new Regex(@"\t|\n|\r");

            // Append a blank line to ensure last passport is returned.
            foreach(var line in File.ReadLines(_inputFilePath).Append("\n"))
            {
                if(string.IsNullOrWhiteSpace(line))
                {
                    yield return passport;
                    passport = string.Empty;
                    continue;
                }

                // Fields are space delimited.
                // To simplify later steps we remove as much "noise" from the string as possible, at this stage.
                passport += $" {newlineAndTabRegEx.Replace(line, "")}";
            }


            Debug.Assert(passport.Trim().Length == 0, "Input contains unread passports");
        }
    }
}
