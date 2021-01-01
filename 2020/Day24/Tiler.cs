using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Tiler
    {
        public void Lay(IEnumerable<TileToken> tiles)
        {
            var blackTiles = new Dictionary<Location, TileToken>();
            var tilesFlipped = 0;

            Console.WriteLine("\nLaying tiles:");
            foreach(var tile in tiles)
                FlipTile(tile);


            Console.WriteLine($"\nStats:\n  Tiles flipped: {tilesFlipped}\n  Black tiles: {blackTiles.Count}\n");
            Flip(blackTiles);
            return;


            void FlipTile(TileToken tile)
            {
                if(blackTiles.ContainsKey(tile.Location))
                    blackTiles.Remove(tile.Location);
                else
                    blackTiles.Add(tile.Location, tile.Flip());

                Console.WriteLine($"  {tile.Location}");
                tilesFlipped++;
            }
        }


        private void Flip(Dictionary<Location, TileToken> blackTiles)
        {
            Console.WriteLine("Simulating Days:");
            for(var day = 1; day <= 100; day++)
            {
                var whiteToFlip = WhiteTilesWithTwoBlackNeighbours(blackTiles).ToList();
                var blackToFlip = BlackTilesWithZeroThreeOrMoreBlackNeigbours(blackTiles).ToList();

                blackToFlip.ForEach(b => blackTiles.Remove(b));
                whiteToFlip.ForEach(w => blackTiles.Add(w, new TileToken(Colour.Black, w, new List<Direction>())));

                Console.WriteLine($"  Day {day}: {blackTiles.Count()}");
            }
        }

        private IEnumerable<Location> WhiteTilesWithTwoBlackNeighbours(Dictionary<Location, TileToken> blackTiles) =>
            from whiteTile in GetWhiteTiles(blackTiles.Keys)
            where CountBlackNeighbours(blackTiles, whiteTile) is 2
            select whiteTile
        ;

        private IEnumerable<Location> BlackTilesWithZeroThreeOrMoreBlackNeigbours(Dictionary<Location, TileToken> blackTiles) =>
            from blackTile in blackTiles
            where CountBlackNeighbours(blackTiles, blackTile.Key) is 0 or > 2
            select blackTile.Key
        ;

        private int CountBlackNeighbours(Dictionary<Location, TileToken> blackTiles, Location location) =>
            GetNeighbours(location).Where(n => blackTiles.ContainsKey(n)).Count()
        ;

        private IEnumerable<Location> GetWhiteTiles(IEnumerable<Location> blackTiles)
        {
            var returnedLocations = new List<Location>();

            // We only need to worry about white tiles that boarder a black tile.
            // The rest of the floor cannot be flipped in the current day.
            foreach(var blackTile in blackTiles)
                foreach(var neighbour in GetNeighbours(blackTile))
                    if( ! blackTiles.Contains(neighbour) )
                        if(IsNewLocation(neighbour))
                            yield return neighbour;


            bool IsNewLocation(Location location)
            {
                if(returnedLocations.Contains(location))
                    return false;


                returnedLocations.Add(location);
                return true;
            }
        }

        private IEnumerable<Location> GetNeighbours(Location location)
        {
            yield return new Location(location.Northing + 1, location.Easting + .5);    // ne
            yield return new Location(location.Northing,     location.Easting +  1);    // e
            yield return new Location(location.Northing - 1, location.Easting + .5);    // se
            yield return new Location(location.Northing - 1, location.Easting - .5);    // sw
            yield return new Location(location.Northing,     location.Easting -  1);    // w
            yield return new Location(location.Northing + 1, location.Easting - .5);    // nw
        }
    }
}
