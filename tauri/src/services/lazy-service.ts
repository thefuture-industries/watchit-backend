class LazyService {
  private readonly ErrorImg: string = "/src/assets/default_image.jpg";
  /**
   * Ленивая загрузка фотографий
   */
  public createImageObserver(
    imgRef: any,
    src: string,
    setPoster: any,
    setLoaded: any,
    onErrorSrc = this.ErrorImg
  ) {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            const img = new Image();
            img.src = src;
            img.onload = () => {
              setPoster(src);
              setLoaded(true);
            };
            img.onerror = () => {
              setPoster(onErrorSrc);
              console.log(src);
              setLoaded(true);
            };

            if (imgRef.current) {
              observer.unobserve(imgRef.current);
            }
          }
        });
      },
      {
        rootMargin: "50px",
        threshold: 0,
      }
    );

    if (imgRef.current) {
      observer.observe(imgRef.current);
    }

    return () => {
      if (imgRef.current) {
        observer.unobserve(imgRef.current);
      }
    };
  }
}

export default new LazyService();
