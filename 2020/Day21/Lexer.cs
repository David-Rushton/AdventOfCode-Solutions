using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Lexer
    {
        string _inputPath = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");

        Dictionary<string, int> _allergenMasterList = new();

        Dictionary<string, int> _ingredientsMasterList = new();

        List<Food> _foods = new();


        public (List<Food> foods, Dictionary<string, int> ingredients, Dictionary<string, int> allergens) GetTokens
        (
            bool useTestDataset
        )
        {
            var input = useTestDataset ? GetTestInput() : GetInput();

            foreach(var line in input)
            {
                var ingredients = new List<string>();
                var allergens = new List<string>();
                var allergensMode = false;
                var words = line.Replace("(", "").Replace(")", "").Replace(",", "");

                foreach(var word in words.Split(' '))
                {
                    if(word == "contains")
                    {
                        allergensMode = true;
                        continue;
                    }

                    if(allergensMode)
                        updateLists(_allergenMasterList, allergens, word);
                    else
                        updateLists(_ingredientsMasterList, ingredients, word);
                }

                _foods.Add(new Food(ingredients, allergens));
            }


            return
            (
                _foods,
                _ingredientsMasterList,
                _allergenMasterList
            );


            void updateLists(Dictionary<string, int> masterList, List<string> list, string item)
            {
                list.Add(item);
                updateMasterList(masterList, item);
            }

            void updateMasterList(Dictionary<string, int> masterList, string item)
            {
                if(masterList.ContainsKey(item))
                    masterList[item]++;
                else
                    masterList.Add(item, 1);
            }
        }


        private string[] GetTestInput() =>
            new []
            {
                "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
                "trh fvjkl sbzzf mxmxvkd (contains dairy)",
                "sqjhc fvjkl (contains soy)",
                "sqjhc mxmxvkd sbzzf (contains fish)"
            }
        ;

        private string[] GetInput() => File.ReadAllLines(_inputPath);
    }
}
