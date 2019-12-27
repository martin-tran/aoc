higherOddSquare :: Int -> Int
higherOddSquare n
    | odd sqrtn = sqrtn * sqrtn
    | otherwise = (sqrtn + 1) * (sqrtn + 1)
    where sqrtn = intSqrt n

intSqrt :: Int -> Int
intSqrt n = ceiling . sqrt $ fromIntegral n

solve :: Int -> Int
solve 1 = 0
solve n = (abs $ ((higherOddSquare n - n) `mod` m) - (m `div` 2)) + (m `div` 2)
  where m = (intSqrt $ higherOddSquare n) - 1

main = do
  n <- getContents
  putStrLn $ (show . solve) $ read 
