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
                batchTotals.Add(ProcessBatchWithPrecedence(tokenQueue));

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


        private long ProcessBatchWithPrecedence(Queue<Token> tokens)
        {
            var buffer = new List<long>();
            var addStack = new Stack<long>();
            var multiplyStack = new Stack<long>();
            var currentOperator = "-";
            long runningTotal = 0;

            while(tokens.Count > 0)
            {
                var token = tokens.Dequeue();
                Console.Write($"{token.Value} ");

                switch (token.Type)
                {
                    case TokenType.Integer:
                        AddToStackOrBuffer(long.Parse(token.Value));
                        break;

                    case TokenType.Multiplication:
                        if(addStack.Count > 0)
                            CollapseToMultipleStack();

                        currentOperator = token.Value;
                        MoveBufferToStack();
                        break;

                    case TokenType.Addition:
                        currentOperator = token.Value;
                        MoveBufferToStack();
                        break;

                    case TokenType.LeftParentheses:
                        AddToStackOrBuffer(ProcessBatchWithPrecedence(tokens));
                        break;

                    case TokenType.RightParentheses:
                        ProcessStacks();
                        return runningTotal;

                    case TokenType.Equals:
                        ProcessStacks();
                        Console.WriteLine($" {runningTotal}");
                        return runningTotal;

                    default:
                        throw new Exception($"Unexpected token: {token}");
                }
            }


            throw new Exception("End of batch not received");



            void AddToStackOrBuffer(long number)
            {
                if(currentOperator == "+")
                    addStack.Push(number);
                else
                    AddToBuffer(number);
            }

            void AddToBuffer(long number) => buffer.Add(number);

            void MoveBufferToStack()
            {
                if(currentOperator == "+")
                    buffer.ForEach(l => addStack.Push(l));
                else
                    buffer.ForEach(l => multiplyStack.Push(l));

                buffer.Clear();
            }

            void CollapseToMultipleStack()
            {
                long subTotal = 0;

                while(addStack.Count > 0)
                    subTotal += addStack.Pop();

                while(multiplyStack.Count > 0)
                    subTotal *= multiplyStack.Pop();

                multiplyStack.Push(subTotal);
            }

            void ProcessStacks()
            {
                MoveBufferToStack();

                while(addStack.Count > 0)
                    runningTotal += addStack.Pop();

                // Avoid multiplying by 0
                if(runningTotal == 0 && multiplyStack.Count > 1)
                    runningTotal = 1;

                while(multiplyStack.Count > 0)
                    runningTotal *= multiplyStack.Pop();
            }
        }
    }
}
