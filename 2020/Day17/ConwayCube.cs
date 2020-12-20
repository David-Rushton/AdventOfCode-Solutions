namespace AoC
{
    public record ConwayCubePosition
    (
        int X,
        int Y,
        int Z,
        int W
    )
    {
        public string Id => string.Format("{0}:{1}:{2}:{3}", this.X, this.Y, this.Z, this.W);
    };


    public record ConwayCube
    (
        ConwayCubePosition Position,
        bool IsActive
    )
    {
        public ConwayCube(int x, int y, int z, int w, bool isActive)
            : this
            (
                new ConwayCubePosition(x, y, z, w),
                isActive
            )
        { }


        // Cube and it's position share a common id.
        // The id is its unique location (x, y and z address).
        public string Id => Position.Id;
    }
}
