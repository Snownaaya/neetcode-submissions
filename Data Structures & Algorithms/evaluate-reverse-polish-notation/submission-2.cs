public class Solution {

    private delegate int _operation(int number1, int number2);

    public int EvalRPN(string[] tokens) {
        Dictionary<string, _operation> operations = new Dictionary<string, _operation>() {
            {"+", (number1, number2) => number1 + number2 },
            {"*", (number1, number2) => number1 * number2 },
            {"-", (number1, number2) => number1 - number2 },
            {"/", (number1, number2) => number1 / number2 }
        };

        Stack<int> stack = new Stack<int>();

        foreach (string token in tokens)
        {
            if (operations.TryGetValue(token, out var operation))
            {
                int number1 = stack.Pop();
                int number2 = stack.Pop();
                stack.Push(operation(number2, number1));
            }
            else
            {
                int result = Convert.ToInt32(token);
                stack.Push(result);
            }
        }

        return stack.Peek();
    }
}
