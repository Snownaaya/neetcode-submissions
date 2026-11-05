public class Solution
{
    public bool IsPalindrome(string s)
    {
        int left = 0;
        int right = s.Length - 1;

        while (left < right)
        {
            while (left < right && char.IsLetterOrDigit(s[left]) == false)
                left++;

            while (left < right && char.IsLetterOrDigit(s[right]) == false)
                right--;

            if (char.ToLower(s[left]) != char.ToLower(s[right]))
                return false;

            left++;
            right--;
        }
        
        return true;
    }
}
