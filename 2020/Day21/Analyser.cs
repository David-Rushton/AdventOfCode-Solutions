using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Analyser
    {
        public void Analyse(List<Food> foods, Dictionary<string, int> ingredients, Dictionary<string, int> allergens)
        {
            var shortlist = GetAllergenShortlist(foods, ingredients.Select(i => i.Key).ToList(), allergens.Select(a => a.Key).ToList());
            var (unmapped, map) = GetAllergenIngredientMap(shortlist);


            if(unmapped.Count() > 0)
            {
                Console.WriteLine($"\nCould not map all allergens.\nUnmapped: { string.Join(", ", unmapped) }");
                Environment.Exit(1);
            }


            var countOfAllergenFreeIngredients = GetCountOfAllergenFreeIngredients(ingredients, map.Select(i => i.Value).ToList());
            Console.WriteLine($"\nAllergen free ingredients: {countOfAllergenFreeIngredients}");
            Console.WriteLine($"Allergens: { string.Join(',', map.OrderBy(a => a.Key).Select(a => a.Value)) }");
        }


        private Dictionary<string, List<string>> GetAllergenShortlist(List<Food> foods, List<string> ingredients, List<string> allergens)
        {
            var shortlist = new Dictionary<string, List<string>>();

            Console.WriteLine("Allergens:");
            foreach(var allergen in allergens)
            {
                var candidateIngredients = foods.Where(f => f.Allergens.Contains(allergen)).Select(f => f.Ingredients);
                var intersectionOfCandidates = candidateIngredients.First();

                foreach(var ingredient in candidateIngredients.Skip(1))
                    intersectionOfCandidates = intersectionOfCandidates.Intersect(ingredient).ToList();

                Console.WriteLine($"  {allergen} shortlist: { string.Join(", ", intersectionOfCandidates) }");
                shortlist.Add(allergen, intersectionOfCandidates);
            }


            return shortlist;
        }

        private (List<string> unmapped, Dictionary<string, string> map) GetAllergenIngredientMap(Dictionary<string, List<string>> shortlist)
        {
            var newMatches = 1;
            var map = new Dictionary<string, string>();

            while(newMatches > 0)
            {
                newMatches = 0;

                foreach(var item in shortlist)
                {
                    var candidates = item.Value.Where(i => ! map.ContainsValue(i));
                    if(candidates.Count() == 1)
                    {
                        map.Add(item.Key, candidates.First());
                        Console.WriteLine($"  Match: {item.Key} = {map[item.Key]}");
                        newMatches++;
                    }
                }
            }


            return
            (
                shortlist.Where(i => ! map.ContainsKey(i.Key)).Select(i => i.Key).ToList(),
                map
            );
        }

        private int GetCountOfAllergenFreeIngredients(Dictionary<string, int> ingredients, List<string> allergens) =>
            (
                from ingredient in ingredients
                where allergens.Contains(ingredient.Key) == false
                select ingredient.Value
            ).Sum()
        ;
    }
}
