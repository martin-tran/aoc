import Data.List

distribute :: [Int] -> Int -> Int -> [Int]
distribute xs 0 i = xs
distribute xs n i = distribute ((take modded xs) ++
                                [(xs !! modded) + 1] ++
                                (drop (if moddedplus == 0 then (length xs) else moddedplus) xs))
                    (n - 1) moddedplus
  where modded = i `mod` (length xs)
        moddedplus = (i + 1) `mod` (length xs)

indexMax :: [Int] -> Int
indexMax xs
  | xs == [] || head xs == maximum xs = 0
  | otherwise = 1 + indexMax (tail xs)

cycleLength :: [Int] -> [[Int]]-> (Int, [[Int]])
cycleLength xs vec
  | findIndex (\x -> x == xs) vec /= Nothing = (0, vec ++ [xs])
  | otherwise = (1 + (fst ans), snd ans)
    where i = (indexMax xs)
          iplus = if 1 + i == length xs then 0 else 1 + i
          xss = (take i xs) ++ [0] ++ (drop (1 + i) xs)
          ans = cycleLength (distribute xss (maximum xs) iplus) (vec ++ [xs])

main = do
  contents <- getContents
  let
    ans = cycleLength [read x :: Int | x <- words contents] []
  putStrLn $ "Part 1: " ++ (show $ fst $ ans)
  putStrLn $ "Part 2: " ++
    (show $ maybe 0 (\x -> (fst ans) - x) (findIndex (\x -> x == last (snd ans)) (snd ans)))
