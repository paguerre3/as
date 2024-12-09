# Handle large files below 2 GB

To manage large files like `mobility_data.parquet` in your cloned repositories using Git Large File Storage (LFS), follow these steps:

---

### **1. Install Git LFS**
Ensure you have Git LFS installed:
- **Linux**: Install via your package manager (e.g., `sudo apt-get install git-lfs`).
- **macOS**: Install via Homebrew (`brew install git-lfs`).
- **Windows**: Download from [Git LFS website](https://git-lfs.github.com/).

After installation, enable LFS in your repository:
```bash
git lfs install
```

---

### **2. Track the File Type**
Set up LFS tracking for `.parquet` files (or specifically for `mobility_data.parquet`):
```bash
# splitted mobility_data.parquet file compressed (.zip)
git lfs track "*mobility_data.z01"
git lfs track "*mobility_data.z02"
git lfs track "*mobility_data.zip"
```
This adds a `.gitattributes` file in the repository root, which ensures that `.parquet` files are managed by LFS.

Commit the `.gitattributes` file:
```bash
git add .gitattributes
git commit -m "Track .parquet files with LFS"
```

---

### **3. Add and Push the File**
Add the large file to your repository and push it to your remote:
```bash
git add internal/modules/assets/mobility_data.parquet
git commit -m "Add mobility_data.parquet using LFS"
git push
```
Git LFS will automatically handle the upload of the file to the LFS storage.

---

### **4. Clone and Download Files**
When someone clones the repository, they need to have LFS installed. Git LFS will automatically download the large file during the cloning process:
```bash
git clone <repository-url>
```

If the file placeholders are downloaded instead of the actual file, they can retrieve the file by running:
```bash
git lfs pull
```

---

### **5. Confirm LFS Functionality**
To verify that LFS is correctly managing your file:
```bash
git lfs ls-files
```
This should list `mobility_data.parquet` as an LFS-managed file.

---

By following these steps, you can efficiently upload and download large files like `mobility_data.parquet` in your Git repositories using LFS.