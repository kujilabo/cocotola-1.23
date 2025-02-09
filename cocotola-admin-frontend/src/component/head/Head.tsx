import { Helmet } from "react-helmet";
type HeadProps = {
  title?: string;
  description?: string;
};

export const Head = ({ title = "", description = "" }: HeadProps = {}) => {
  return (
    <Helmet>
      <meta charSet="utf-8" />
      <title>{title}</title>
      {/* <link rel="canonical" href="http://mysite.com/example" /> */}
    </Helmet>
  );
};
