using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static System.Console;
using static System.Math;

namespace Program
{
    class Program
    {
        static void Main(string[] args)
        {
            Do();
        }
        static void Do()
        {
            int[] rc = ReadLine().Split(' ').Select(int.Parse).ToArray();
            int r = rc[0];
            int c = rc[1];
            var l = new List<string>(r);
            for(int i = 0 ; i < c ; i++)
            {
                l.Add(ReadLine());
            }
        }
    }
}

/* 
// 読み込み関連
// 文字列読み込み
var str = ReadLine();
// 文字列配列読み込み
var strs = ReadLine().Split(' ');
// 数字読み込み
int N = int.Parse(ReadLine());
// 数値配列読み込み
int[] ia = ReadLine().Split(' ').Select(int.Parse).ToArray();

// 変換
var i = int.Parse(str);
var l = long.Parse(str);

// split
var strs = str.Split(' ');

// join
string.join(",",strs);

// 書き出し
WriteLine($"{str}");

// Linq
str
.GroupBy(x => x)
.Select(x => new { n = x.Key,c = (int)x.Key, count = x.Count() })
.OrderBy(x => x.c).ToList();
*/

//10
/*
            var str = ReadLine();
            var i = 0;
            List<char> list = new List<char>();
            foreach(var c in str)
            {
                i = c == '(' ? i + 1 : i -1;
                if(i < 0)
                {
                    list.Insert(0,'(');
                    list.Add(c);
                    i = 0;
                }else{
                    list.Add(c);
                }
            }
            while(i > 0)
            {
                list.Add(')');
                i--;
            }
            WriteLine(new String(list.ToArray()));
 */
