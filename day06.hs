import qualified Data.Set

distribute :: [Int] -> Int -> Int -> [Int]
distribute xs 0 i = xs
distribute xs n i = distribute ((take modded xs) ++
                                [(xs !! modded) + 1] ++
                                (drop (if moddedplus == 0 then (length xs) else moddedplus) xs))
                    (n - 1) moddedplus
  where modded = i `mod` (length xs)
        moddedplus = (i + 1) `mod` (length xs)
