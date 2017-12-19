countSteps1 :: [Int] -> Int -> Int
countSteps1 xs i
  | i + (xs !! i) >= length xs = 1
  | otherwise = 1 + countSteps1 ((take i xs) ++ [(xs !! i) + 1] ++ (drop (i + 1) xs)) (i + (xs !! i))
  
countSteps2 :: [Int] -> Int -> Int
countSteps2 xs i
  | i + (xs !! i) < 0 || i + (xs !! i) >= length xs = 1
  | xs !! i >= 3 = 1 + countSteps2
    ((take i xs) ++ [(xs !! i) - 1] ++ (drop (i + 1) xs))
    (i + (xs !! i))
  | otherwise = 1 + countSteps2 ((take i xs) ++ [(xs !! i) + 1] ++ (drop (i + 1) xs)) (i + (xs !! i))

main = do
  contents <- getContents
  putStrLn $ "Part 1: " ++ (show $ countSteps1 [read x :: Int | x <- lines contents] 0)
  putStrLn $ "Part 2: " ++ (show $ countSteps2 [read x :: Int | x <- lines contents] 0)
