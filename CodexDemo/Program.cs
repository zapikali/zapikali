using System;

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Расчёт среднего возраста команды.");

        int[] ages = new int[3];

        for (int i = 0; i < ages.Length; i++)
        {
            Console.Write($"Введите возраст #{i + 1}: ");
            var ageText = Console.ReadLine();

            if (!int.TryParse(ageText, out var age))
            {
                Console.WriteLine("Некорректный возраст.");
                return;
            }

            ages[i] = age;
        }

        double averageAge = (ages[0] + ages[1] + ages[2]) / 3.0;
        Console.WriteLine($"Средний возраст: {averageAge:F2}");

        if (averageAge < 18)
        {
            Console.WriteLine("Команда юных разработчиков");
        }
        else
        {
            Console.WriteLine("Команда взрослых разработчиков");
        }
    }
}
