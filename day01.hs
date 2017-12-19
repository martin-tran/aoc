-- Advent of Code 2017 - Reverse Captcha
import Data.Char

sumAhead :: [Int] -> [Int] -> Int
sumAhead [] _ = 0
sumAhead (x:xs) (y:ys)
  | x == y = x + sumAhead xs ys
  | otherwise = sumAhead xs ys

shiftN :: Int -> [Int] -> [Int]
shiftN n xs = (drop nmod xs) ++ (take nmod xs)
  where nmod = n `mod` (length xs)

main = do
  inString <- getLine
  let intList = [digitToInt x | x <- inString]
      half = length intList `div` 2
  putStrLn $ "Part 1: " ++ (show $ sumAhead intList $ shiftN 1 intList)
  putStrLn $ "Part 2: " ++ (show $ sumAhead intList $ shiftN half intList)
