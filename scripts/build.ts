import AdmZip from 'adm-zip';

(async () => {
  const zip = new AdmZip();
  zip.addLocalFile('./packages/frontend/dist/index.html');
  await zip.addLocalFolderPromise('./packages/frontend/dist/assets', {});
  await zip.writeZipPromise('./packages/backend/assets.zip');
})();
