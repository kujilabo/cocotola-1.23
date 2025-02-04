type HeadProps = {
  title?: string;
  description?: string;
};

export const Head = ({ title = '', description = '' }: HeadProps = {}) => {
  return (
    <>
      {title} {description}
    </>
    // <Helmet
    //   title={title ? `${title} | Bulletproof React` : undefined}
    //   defaultTitle="Bulletproof React"
    // >
    //   <meta name="description" content={description} />
    // </Helmet>
  );
};