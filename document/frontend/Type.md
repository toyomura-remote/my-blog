
src/types/post.tsを作成してinterfaceを定義
```
export interface Post {
  id: number;
  title: string;
  content: string;
  // 他に createdAt などがあれば適宜追加
}
```

他の型をインポートするときは`type`をつける
```
import type { Post } from "./post";
```

Ginを使用する場合data{ hoge: hoge}でデータが返ってくるので、それの受け皿を用意する
```
export interface PostResponse {
    data: Post[];
} 
```