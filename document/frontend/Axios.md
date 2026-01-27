**パッケージのインストール**
npm install axios

**まず共有のマシンを作る**

`src/api/cliant.ts`
```
export const apiCliant = axios.create({
    baseURL: API_BASE_URL,
    timeout: 5000,
    headers:{
        'Content-Type': 'application/json',
    }
})
```

**使わない場合（非効率）:**

TypeScript

```
axios.get('http://localhost:8080/posts');
axios.post('http://localhost:8080/posts', data); // 毎回URLを書くのが面倒
```

**使った場合（効率的）:**

TypeScript

```
apiClient.get('/posts'); // 自動的に http://localhost:8080/posts になる
```