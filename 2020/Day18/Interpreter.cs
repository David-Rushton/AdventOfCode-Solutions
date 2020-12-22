using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.IO;


namespace AoC
{
    public class Interpreter
    {
        public void Calculate(IEnumerable<Token> tokens)
        {
            var tokenQueue = new Queue<Token>(tokens);
            var batchTotals = new List<long>();


            while(tokenQueue.Count > 0)
                batchTotals.Add(ProcessBatch(tokenQueue));

            Console.WriteLine($"\nTotal: {batchTotals.Sum()}");
        }


        private long ProcessBatch(Queue<Token> tokens)
        {
            long runningTotal = 0;
            var currentOperator = "+";
            var expression = string.Empty;

            while(tokens.Count > 0)
            {
                var token = tokens.Dequeue();
                Console.Write($"{token.Value} ");

                switch (token.Type)
                {
                    case TokenType.Integer:
                        runningTotal = EvaluateExpression(runningTotal, currentOperator, long.Parse(token.Value));
                        break;

                    case TokenType.Addition:
                    case TokenType.Multiplication:
                        currentOperator = token.Value;
                        break;

                    case TokenType.LeftParentheses:
                        runningTotal = EvaluateExpression(runningTotal, currentOperator, ProcessBatch(tokens));
                        break;

                    case TokenType.RightParentheses:
                        return runningTotal;

                    case TokenType.Equals:
                        Console.WriteLine(runningTotal);
                        return runningTotal;

                    default:
                        throw new Exception($"Unexpected token: {token}");
                }
            }


            throw new Exception("End of batch not received");


            long EvaluateExpression(long left, string operation, long right) =>
                operation == "+" ? left + right : left * right
            ;
        }
    }
}
