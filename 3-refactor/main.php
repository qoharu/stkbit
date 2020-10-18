<?
echo findFirstStringInBracket($argv[1]);

function findFirstStringInBracket($str){
      // 1. Trim string before first starting bracket `(`
      $trimmedFirstBracketedWord = strstr($str, '(');

      // 2. Remove multiple starting bracket `(`
      $removedMultipleStartingBrackets = ltrim($trimmedFirstBracketedWord, '(');

      // 3. Remove string after ending bracket `)` and return the result
      return strstr($trimFirstBracketedWord, ')', true);
}