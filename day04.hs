import qualified Data.List
import qualified Data.Set

validPhrase :: [String] -> Bool
validPhrase [] = True
validPhrase xs = (length xs) == (Data.Set.size $ Data.Set.fromList xs)

main = do
  contents <- getContents
  putStrLn $ "Part 1: " ++ (show $ length [x | x <- [words s | s <- lines contents], validPhrase x])
  putStrLn $ "Part 2: " ++ (show $ length [x | x <- [map Data.List.sort $ words s |
                                                     s <- lines contents], validPhrase x])
