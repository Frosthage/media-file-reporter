using System;
using Snapper;
using Snapper.Attributes;
using Xunit;

namespace TestProject1
{
    public class UnitTest1
    {
        [Fact]
        public void Test1()
        {
            new Apa().ShouldMatchSnapshot();
        }
    }

    class Apa
    {
        public int X { get; set; } = 667;
        public string Yyyy { get; set; } = "666";
        public double ZzzzZzzz { get; set; } = 6.67;
        public ApaEnum WwwWWWWwwWwwwW { get; set; } = ApaEnum.Cepa;
    }

    enum ApaEnum
    {
        Apa,
        Bepa,
        Cepa,
    }
    
    
}