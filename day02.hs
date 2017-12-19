-- Advent of Code 2017 - Corruption Checksum
bigDif :: [Int] -> Int
bigDif xs = maximum xs - minimum xs

evenDiv :: [Int] -> Int
evenDiv xs = case [x | x <- drop 1 xs, ((head xs) `mod` x) == 0] of
  [] -> evenDiv $ drop 1 xs ++ [head xs]
  [y] -> head xs `div` y

main = do
  contents <- getContents
  let matrix = [[read x :: Int | x <- words xs] | xs <- lines contents]
  putStrLn $ "Part 1: " ++ (show $ sum $ map bigDif matrix)
  putStrLn $ "Part 2: " ++ (show $ sum $ map evenDiv matrix)
