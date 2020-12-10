using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace Day09
{
    class Program
    {
        static void Main(string[] args)
        {
            var testMode = args.Contains("--test");
            var inputPath = Path.Join(Directory.GetCurrentDirectory(), testMode ? "Test-Input.txt" : "Input.txt");

            if(args.Length > 0)
                if(long.TryParse(args[0], out var target))
                    StarTwo(testMode, inputPath, target);

            StarOne(testMode, inputPath);
        }

        static void StarTwo(bool testMode, string inputPath, long target)
        {
            long total = 0;
            var queue = new Queue<long>();

            foreach(var value in GetInput(inputPath))
            {
                total += value;
                queue.Enqueue(value);

                while(total > target)
                {
                    total -= queue.Dequeue();
                }


                if(total == target)
                    ReportAndExit();
            }


            Console.WriteLine("Target not found\n:(");
            Environment.Exit(1);


            void ReportAndExit()
            {
                if(total == target && queue.Count > 1)
                {
                    var smallest = queue.Min();
                    var largest = queue.Max();

                    Console.WriteLine($"Target found!\n  Contiguous set:\n    Smallest number:{smallest}\n    Largest number: {largest}\n    Contiguous set size: {queue.Count}");
                    Console.WriteLine($"Result: {smallest + largest}");
                    Environment.Exit(0);
                }
            }
        }

        static void StarOne(bool testMode, string inputPath)
        {
            var queueSize = testMode ? 5 : 25;
            var input = GetInput(inputPath);
            var queue = new Queue<long>(input.Take(queueSize));

            foreach(var value in input.Skip(queueSize))
            {
                for(var o = 0; o < queue.Count; o++)
                {
                    for(var i = 1; i < queue.Count; i++)
                    {
                        if(queue.ElementAt(o) + queue.ElementAt(i) == value)
                            goto ValidValue;
                    }
                }

                // This section will only execute if value is the not the product of any 2 items in the queue
                Console.WriteLine($"\nResult: {value}");
                Environment.Exit(0);


                ValidValue:
                queue.Dequeue();
                queue.Enqueue(value);
            }


            Console.WriteLine("Program failed");
            Environment.Exit(1);
        }


        static IEnumerable<Int64> GetInput(string path)
        {
            foreach(var line in File.ReadLines(path))
            {
                yield return long.Parse(line);
            }
        }

    }
}
