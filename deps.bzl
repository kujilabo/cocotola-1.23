load("@gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "cat_dario_mergo",
        importpath = "dario.cat/mergo",
        sum = "h1:AGCNq9Evsj31mOgNPcLyXc+4PNABt905YmuqPYYpBWk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
        version = "v0.0.1-2020.1.4",
    )
    go_repository(
        name = "com_github_99designs_go_keychain",
        importpath = "github.com/99designs/go-keychain",
        sum = "h1:/vQbFIOMbk2FiG/kXiLl8BRyzTWDw7gX/Hz7Dd5eDMs=",
        version = "v0.0.0-20191008050251-8e49817e8af4",
    )
    go_repository(
        name = "com_github_99designs_keyring",
        importpath = "github.com/99designs/keyring",
        sum = "h1:tYLp1ULvO7i3fI5vE21ReQuj99QFSs7lGm0xWyJo87o=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_adalogics_go_fuzz_headers",
        importpath = "github.com/AdaLogics/go-fuzz-headers",
        sum = "h1:He8afgbRMd7mFxO99hRNu+6tazq8nFF9lIwo9JFroBk=",
        version = "v0.0.0-20240806141605-e8a1dd7889d6",
    )

    go_repository(
        name = "com_github_agiledragon_gomonkey_v2",
        importpath = "github.com/agiledragon/gomonkey/v2",
        sum = "h1:QJWqpdEhGV/JJy70sZ/LDnhbSlMrqHAWHcNOjz1kyuI=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_antihax_optional",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v10",
        importpath = "github.com/apache/arrow/go/v10",
        sum = "h1:n9dERvixoC/1JjDmBcs9FPaEryoANa2sCgVFo6ez9cI=",
        version = "v10.0.1",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v15",
        importpath = "github.com/apache/arrow/go/v15",
        sum = "h1:60IliRbiyTWCWjERBCkO1W4Qun9svcYoZrSLcyOsMLE=",
        version = "v15.0.2",
    )
    go_repository(
        name = "com_github_apache_thrift",
        importpath = "github.com/apache/thrift",
        sum = "h1:qEy6UW60iVOlUy+b9ZR0d5WzUWYGOo4HfopoyBaNmoY=",
        version = "v0.16.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:yNldzF5kzLBRvKlKz1S0bkvc2+04R1kt13KfBWQBfFA=",
        version = "v1.49.6",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:M1fj4FE2lB4NzRb9Y0xdWsn2P0+2UHVxwKyOa4YJNjk=",
        version = "v1.16.16",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:tcFliCWne+zOuUfKNRn8JdFBuWPDuISDH08wD2ULkhk=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:9+ZhlDY7N9dPnUmf7CDfW9In4sW5Ff3bh7oy4DzS1IE=",
        version = "v1.12.20",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_s3_manager",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/s3/manager",
        sum = "h1:fAoVmNGhir6BR+RU0/EI+6+D7abM+MCwWf8v4ip5jNI=",
        version = "v1.11.33",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:s4g/wnzMf+qepSNgTvaQQHNxyMLKSawNhKCPNy++2xY=",
        version = "v1.1.23",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:/K482T5A3623WJgWT8w1yRAFK4RzGzEl7y39yhtn9eA=",
        version = "v2.4.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:ZSIPAkAsCCjYrhqfw2+lNzWDzxzHXEckFkTePL5RSWQ=",
        version = "v1.0.14",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:Lh1AShsuIJTwMkoxVCAYPJgNG5H+eN6SmoUn8nOZ5wE=",
        version = "v1.9.9",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:BBYoNQt2kUZUUK4bIPsKrCcjVPUMNsgQpNAwhznK/zo=",
        version = "v1.1.18",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:Jrd/oMh0PKQc6+BowB+pLEwLIgaQF29eYbe7E1Av9Ug=",
        version = "v1.9.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:HfVVR1vItaG6le+Bpw6P4midjBDMKnjMyZnw9MXYUcE=",
        version = "v1.13.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:3/gm/JTX9bX8CpzTgIlrtYpB3EVBDxyg/GY/QdcIEZw=",
        version = "v1.27.11",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:l7LYxGuzK6/K+NzJ2mC+VvLUbae0sL3bXU//04MkmnA=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
        sum = "h1:E+OJmp2tPvt1W+amx48v1eqbjDYsgN+RzP4q16yV5eM=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azidentity",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azidentity",
        sum = "h1:U2rTu3Ef+7w9FHKIAXM6ZyqF3UOWJZ12zIm8zECAFfg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_internal",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
        sum = "h1:jBQA3cKT4L2rWMpgE7Yt3Hwh2aUj8KXjIGLxjHeYNNo=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_security_keyvault_azkeys",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys",
        sum = "h1:MyVTgWR8qd/Jw1Le0NZebGBUCLbtak3bJ3z1OlqZBpw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_security_keyvault_internal",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal",
        sum = "h1:D3occbWoio4EBLkbkevetNMAVX197GkzbUMtqjGWn80=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_storage_azblob",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob",
        sum = "h1:u/LLAOFgsMv7HmNL4Qufg58y+qElGOt5qv0z1mURkRY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:L/gRVlceqvL25UVaW/CKtUDjefjrs0SPonmDGUVOYP0=",
        version = "v0.0.0-20230124172434-306776ec8161",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:P8An8Z9rH1ldbOLdFpxYorgOt2sywL9V24dAwWHPuGc=",
        version = "v0.9.16",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_azuread_microsoft_authentication_library_for_go",
        importpath = "github.com/AzureAD/microsoft-authentication-library-for-go",
        sum = "h1:XHOnouVk1mxXfQidrMEnLlPk9UMeRtyBTnEFtxkV0kU=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_bazelbuild_rules_go",
        importpath = "github.com/bazelbuild/rules_go",
        sum = "h1:5vCbuvy8Q11g41lseGJDc5vxhDjJtfxr6nM/IC4VmqM=",
        version = "v0.49.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bmatcuk_doublestar_v4",
        importpath = "github.com/bmatcuk/doublestar/v4",
        sum = "h1:54Bopc5c2cAvhLRAzqOGCYHYyhcDHsFF4wWIR5wKP38=",
        version = "v4.8.1",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_bytedance_sonic",
        importpath = "github.com/bytedance/sonic",
        sum = "h1:Od1BvK55NnewtGaJsTDeAOSnLVO2BTSLOe0+ooKokmQ=",
        version = "v1.12.9",
    )
    go_repository(
        name = "com_github_bytedance_sonic_loader",
        importpath = "github.com/bytedance/sonic/loader",
        sum = "h1:yctD0Q3v2NOGfSWPLPvG2ggA2kV6TS6s4wioyEqssH0=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_casbin_casbin_v2",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:dHElatNXNrr8XcseUov0ZSiWjauwmZZE6YMV3eU1yic=",
        version = "v2.103.0",
    )
    go_repository(
        name = "com_github_casbin_gorm_adapter_v3",
        importpath = "github.com/casbin/gorm-adapter/v3",
        sum = "h1:Au+IOILBIE9clox5BJhI2nA3p9t7Ep1ePlupdGbGfus=",
        version = "v3.32.0",
    )
    go_repository(
        name = "com_github_casbin_govaluate",
        importpath = "github.com/casbin/govaluate",
        sum = "h1:VA0eSY0M2lA86dYd5kPPuNZMUD9QkWnOCnavGrw9myc=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:MyRJ/UdXutAwSAT+s3wNd7MfTIcy71VQueUuFK343L8=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_chenzhuoyu_base64x",
        importpath = "github.com/chenzhuoyu/base64x",
        sum = "h1:77cEq6EriyTZ0g/qfRdp61a3Uu/AWrgIq2s0ClJV1g0=",
        version = "v0.0.0-20230717121745-296ad89f973d",
    )
    go_repository(
        name = "com_github_chenzhuoyu_iasm",
        importpath = "github.com/chenzhuoyu/iasm",
        sum = "h1:tUHQJXo3NhBqw6s33wkGn9SP3bvrWLdlVIJ3hQBL7P0=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_chromedp_cdproto",
        importpath = "github.com/chromedp/cdproto",
        sum = "h1:aPflPkRFkVwbW6dmcVqfgwp1i+UWGFH6VgR1Jim5Ygc=",
        version = "v0.0.0-20230802225258-3cf4e6d46a89",
    )
    go_repository(
        name = "com_github_chromedp_chromedp",
        importpath = "github.com/chromedp/chromedp",
        sum = "h1:dKtNz4kApb06KuSXoTQIyUC2TrA0fhGDwNZf3bcgfKw=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_chromedp_sysutil",
        importpath = "github.com/chromedp/sysutil",
        sum = "h1:+ZxhTpfpZlmchB58ih/LBHX52ky7w2VhQVKQMucy3Ic=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        importpath = "github.com/chzyer/logex",
        sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
        version = "v1.1.10",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        importpath = "github.com/chzyer/readline",
        sum = "h1:upd/6fQk4src78LMRzh5vItIt361/o4uq553V8B5sGI=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_chzyer_test",
        importpath = "github.com/chzyer/test",
        sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
        version = "v0.0.0-20180213035817-a1ea475d72b1",
    )
    go_repository(
        name = "com_github_clickhouse_clickhouse_go",
        importpath = "github.com/ClickHouse/clickhouse-go",
        sum = "h1:iAFMa2UrQdR5bHJ2/yaSLffZkxpcOYQMCUuKeNXGdqc=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudflare_golz4",
        importpath = "github.com/cloudflare/golz4",
        sum = "h1:F1EaeKL/ta07PY/k9Os/UFtwERei2/XzGemhpGnBKNg=",
        version = "v0.0.0-20150217214814-ef862a3cdc58",
    )
    go_repository(
        name = "com_github_cloudwego_base64x",
        importpath = "github.com/cloudwego/base64x",
        sum = "h1:XPciSp1xaq2VCSt6lF0phncD4koWyULpl5bUxbfCyP4=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_cloudwego_iasm",
        importpath = "github.com/cloudwego/iasm",
        sum = "h1:1KNIy1I1H9hNNFEEH3DVnI4UujN+1zjpuk6gwHLTssg=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:Om6kYQYDUk5wWbT0t0q6pvyM49i9XZAv9dDrkDA7gjk=",
        version = "v0.0.0-20250121191232-2f005788dc42",
    )
    go_repository(
        name = "com_github_cockroachdb_cockroach_go_v2",
        importpath = "github.com/cockroachdb/cockroach-go/v2",
        sum = "h1:3XzfSMuUT0wBe1a3o5C0eOTcArhmmFAg2Jzh/7hhKqo=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        importpath = "github.com/containerd/containerd",
        sum = "h1:jqjZTQNfXGoEaZdW1WwPU0RqSn1Bm2Ay/KJPUuO8nao=",
        version = "v1.7.18",
    )
    go_repository(
        name = "com_github_containerd_log",
        importpath = "github.com/containerd/log",
        sum = "h1:TCJt7ioM2cr/tfR8GPbGf9/VRAX8D2B4PjzCpfX540I=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_containerd_platforms",
        importpath = "github.com/containerd/platforms",
        sum = "h1:zvwtM3rz2YHPQsF2CHYM8+KtB5dvhISiXh5ZpSBQv6A=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_cpuguy83_dockercfg",
        importpath = "github.com/cpuguy83/dockercfg",
        sum = "h1:DlJTyZGBDlXqUZ2Dk2Q3xHs/FtnooJJVaad2S9GKorA=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_cznic_mathutil",
        importpath = "github.com/cznic/mathutil",
        sum = "h1:XNT/Zf5l++1Pyg08/HV04ppB0gKxAqtZQBRYiYrUuYk=",
        version = "v0.0.0-20180504122225-ca4c9f2c1369",
    )
    go_repository(
        name = "com_github_danieljoos_wincred",
        importpath = "github.com/danieljoos/wincred",
        sum = "h1:QLdCxFs1/Yl4zduvBdcHB8goaYk9RARS2SgLLRuAyr0=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_dhui_dktest",
        importpath = "github.com/dhui/dktest",
        sum = "h1:+I4s6JRE1yGuqflzwqG+aIaMdgXIorCf5P98JnaAWa8=",
        version = "v0.4.4",
    )
    go_repository(
        name = "com_github_distribution_reference",
        importpath = "github.com/distribution/reference",
        sum = "h1:0IXCQ5g4/QMHHkarYzh5l+u8T3t73zM5QvfrDyIgxBk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_dnaeon_go_vcr",
        importpath = "github.com/dnaeon/go-vcr",
        sum = "h1:zHCHvJYTMh1N7xnV7zf1m1GPBF9Ad0Jk/whtQ1663qI=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:Rk9nIVdfH3+Vz4cyI/uhbINhEZ/oLmc+CBXmH6fbNk4=",
        version = "v27.2.0+incompatible",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:USnMq7hx7gwdVZq1L49hLXaFtUdTADjXGp+uj1Br63c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_dvsekhvalnov_jose2go",
        importpath = "github.com/dvsekhvalnov/jose2go",
        sum = "h1:Y9gnSnP4qEI0+/uQkHvFXeD2PLPJeXEL+ySMEA2EjTY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:aaQcKT9WumO6JEJcRyTqFVq4XUZiUcKR2/GI31TOcz8=",
        version = "v0.0.0-20170320065105-0bce6a688712",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:zEqyPVyku6IvWCFwux4x9RxkLOMUL+1vC9xUFv5l2/M=",
        version = "v0.13.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_envoy",
        importpath = "github.com/envoyproxy/go-control-plane/envoy",
        sum = "h1:jb83lalDRZSpPWW2Z7Mck/8kXZ5CQAFYVjQcdVIr83A=",
        version = "v1.32.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_ratelimit",
        importpath = "github.com/envoyproxy/go-control-plane/ratelimit",
        sum = "h1:/G9QYbddjL25KvtKTv3an9lx6VBE2cnb8wp1vEGNYGI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:DEo3O99U8j4hBFwbJfrz9VtgcDfUKS7KJ7spH3d86P8=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_felixge_fgprof",
        importpath = "github.com/felixge/fgprof",
        sum = "h1:VvyZxILNuCiUCSXtPtYmmtGvb65nqXh2QFWc0Wpf2/g=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:/l4kBbb4/vGSsdtB5nUe8L7B9mImVMaBPw9L/0TBHU8=",
        version = "v3.2.5+incompatible",
    )
    go_repository(
        name = "com_github_fsouza_fake_gcs_server",
        importpath = "github.com/fsouza/fake-gcs-server",
        sum = "h1:OeH75kBZcZa3ZE+zz/mFdJ2btt9FgqfjI7gIh9+5fvk=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_gabriel_vasile_mimetype",
        importpath = "github.com/gabriel-vasile/mimetype",
        sum = "h1:FfZ3gj38NjllZIeJAmMhr+qKL8Wu+nOoI3GqacKw1NM=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_gin_contrib_cors",
        importpath = "github.com/gin-contrib/cors",
        sum = "h1:hV+a5xp8hwJoTw7OY+a70FsL8JkVVFTXw9EcfrYUdns=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_github_gin_contrib_sse",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:y3bT1mUWUxDpW4JLQg/HnTqV4rozuW4tC9eFKTxYI9E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:nTuyha1TYqgedzytsKYqna+DfLos46nTv2ygFy86HFU=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_glebarez_go_sqlite",
        importpath = "github.com/glebarez/go-sqlite",
        sum = "h1:uAcMJhaA6r3LHMTFgP0SifzgXg46yJkgxqyuyec+ruQ=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_github_glebarez_sqlite",
        importpath = "github.com/glebarez/sqlite",
        sum = "h1:wSG0irqzP6VurnMEpFGer5Li19RpIRi2qvQz++w0GMw=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_go_gl_glfw",
        importpath = "github.com/go-gl/glfw",
        sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version = "v0.0.0-20190409004039-e6da0acd62b1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )

    go_repository(
        name = "com_github_go_kit_log",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:6pFjapn8bFcIbiKo3XT4j/BhANplGihG6tvd+8rYgrY=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_ole_go_ole",
        importpath = "github.com/go-ole/go-ole",
        sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
        version = "v1.2.6",
    )

    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:MYlu0sBgChmCfJxxUKZ8g1cPWFOB37YSZqewK7OKeyA=",
        version = "v0.20.0",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:1Rlu/ZrOCCob0n+JKKJAWhNWMPW8bOZRg8FJaY+0SKI=",
        version = "v0.20.7",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=",
        version = "v0.22.3",
    )
    go_repository(
        name = "com_github_go_playground_assert_v2",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        importpath = "github.com/go-playground/locales",
        sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_go_playground_validator_v10",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:5Dh7cjvzR7BRZadnsVOzPhWsrwUr0nmsZJxEAnFLNO8=",
        version = "v10.25.0",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:Y0zIbQXhQKmQgTp44Y1dp3wTXcn804QoTptLZT1vtvo=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_gobuffalo_here",
        importpath = "github.com/gobuffalo/here",
        sum = "h1:hYrd0a6gDmWxBM4TnrGw8mQg24iSVoIkHEk7FodQcBI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:exrUm0f4YX0L7EBwZHuCF4GDp8aJfVeBrlLQrs6NqWU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        importpath = "github.com/gobwas/pool",
        sum = "h1:xfeeEhW7pwmX8nuLVlqbzVc7udMDrwetjEv+TZIz1og=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        importpath = "github.com/gobwas/ws",
        sum = "h1:F2aeBZrm2NDsc7vbovKrWSogd4wvfAxg0FQ89/iqOTk=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_goccy_go_json",
        importpath = "github.com/goccy/go-json",
        sum = "h1:Fq85nIqj+gXn/S5ahsiTlK3TmC85qgirsdTP/+DeaC4=",
        version = "v0.10.5",
    )
    go_repository(
        name = "com_github_gocql_gocql",
        importpath = "github.com/gocql/gocql",
        sum = "h1:N/MD/sr6o61X+iZBAT2qEUF023s4KbA8RWfKzl0L6MQ=",
        version = "v0.0.0-20210515062232-b7ef815b4556",
    )
    go_repository(
        name = "com_github_godbus_dbus",
        importpath = "github.com/godbus/dbus",
        sum = "h1:ZpnhV/YsD2/4cESfV5+Hoeu/iUR3ruzNvZ+yQfO03a0=",
        version = "v0.0.0-20190726142602-4481cbc300e2",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:oDTdz9f5VGVVNGu/Q7UXKWYsD0873HXLHdJUNBsSEKM=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:JdqV9zKUdtaa9gdPlywC3aeoEsR681PlKC+4F5gQgeo=",
        version = "v4.5.1",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v5",
        importpath = "github.com/golang-jwt/jwt/v5",
        sum = "h1:OuVbFODueb089Lh128TAcimifWaLhJwVflnrgM17wHk=",
        version = "v5.2.1",
    )
    go_repository(
        name = "com_github_golang_migrate_migrate_v4",
        importpath = "github.com/golang-migrate/migrate/v4",
        sum = "h1:2VSCMz7x7mjyTXx3m2zPokOY82LTRgxK1yQYKo6wWQ8=",
        version = "v4.18.2",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:YojYx61/OLFsiv6Rw1Z96LpldJIy31o+UHmwAUMJ6/U=",
        version = "v1.7.0-rc.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:au07oEsX2xN0ktxqI+Sida1w446QrXBRJ0nee3SNZlA=",
        version = "v0.0.0-20220223132316-b832511892a9",
    )
    go_repository(
        name = "com_github_golang_sql_sqlexp",
        importpath = "github.com/golang-sql/sqlexp",
        sum = "h1:ZCD6MBpcuOVfGVqsEmY5/4FtYiKz6tSyUv9LPEDei6A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:CVpQJjYgC4VbzxeGVHfvZrv1ctoYCAI8vbl07Fcxlyg=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_google_flatbuffers",
        importpath = "github.com/google/flatbuffers",
        sum = "h1:M9dgRyhJemaM4Sw8+66GHBu8ioaQmyPLg1b8VwK5WJg=",
        version = "v23.5.26+incompatible",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:wk8382ETsv4JYUZwIsn6YpYiWiBsYLSJiTsyBybVuN8=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_google_go_github_v39",
        importpath = "github.com/google/go-github/v39",
        sum = "h1:rNNM311XtPOz5rDdsJXAp2o8F67X9FnROXTvto3aSnQ=",
        version = "v39.2.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:PVRnTgtArZ3QQqTGtbtjtnIkzl2iY2kt24yqbrf7td8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_google_martian",
        importpath = "github.com/google/martian",
        sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:DIhPTQrbPkgs2yJYdXU/eNACCG5DVQjySNRNlflZ9Fc=",
        version = "v3.3.3",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:gbpYu9NMq8jhDVbvlGkMFWCjLFlqqEZjEmObmhUy6Vo=",
        version = "v0.0.0-20240409012703-83162a5b38cd",
    )
    go_repository(
        name = "com_github_google_renameio",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:LGD7gtMgezd8a/Xak7mEWL0PjoTQFvpRudN895yqKW0=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_googleapis_cloud_bigtable_clients_test",
        importpath = "github.com/googleapis/cloud-bigtable-clients-test",
        sum = "h1:afMKTvA/jc6jSTMkeHBZGFDTt8Cc+kb1ATFzqMK85hw=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:VgzTY2jogw3xt39CusEnFJWm7rlsq5yL5q9XdLOuP5g=",
        version = "v0.3.5",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:hb0FFeiPaQskmvakKu5EbCbpntQn48jyHuvrkurSS/Q=",
        version = "v2.14.1",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
        version = "v0.0.0-20200911160855-bcd43fbb19e8",
    )
    go_repository(
        name = "com_github_googlecloudplatform_grpc_gcp_go_grpcgcp",
        importpath = "github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp",
        sum = "h1:DBjmt6/otSdULyJdVg2BlG0qGZO5tKL4VzOs0jpvw5Q=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_detectors_gcp",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp",
        sum = "h1:ErKg/3iS1AKcTkf3yixlZ54f9U1rljCkQyEXWUnIUxc=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_metric",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric",
        sum = "h1:fYE9p3esPxA/C0rQ0AHhP0drtPXDRhaWiwg1DPqO7IU=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_trace",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace",
        sum = "h1:Jtr816GUk6+I2ox9L/v+VcOwN6IyGOEDTSNHfD6m9sY=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_cloudmock",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/cloudmock",
        sum = "h1:OqVGm6Ei3x5+yZmSJG1Mh2NwHvpVmZ08CB5qJhT9Nuk=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_resourcemapping",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping",
        sum = "h1:6/0iUd0xrnX7qt+mLNRwg5c0PGv8wpE8K90ryANQwMI=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:0QniY0USkHQ1RGCLfKxeNHK9bkDHGRYGNDFBCS+YARg=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:VuZ8uybHlWmqV03+zRzdwKL4tUnIp1MAQtp1mIFE1bc=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:e9Rjr40Z98/clHv5Yg79Is0NtosR5LXRvdr7o/6NwbA=",
        version = "v2.26.1",
    )
    go_repository(
        name = "com_github_gsterjov_go_libsecret",
        importpath = "github.com/gsterjov/go-libsecret",
        sum = "h1:6rhixN/i8ZofjG1Y75iExal34USq5p+wiN1tpie8IrU=",
        version = "v0.0.0-20161001094733-a6f4afe4910c",
    )
    go_repository(
        name = "com_github_hailocab_go_hostpool",
        importpath = "github.com/hailocab/go-hostpool",
        sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
        version = "v0.0.0-20160125115350-e80d13ce29ed",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
        version = "v0.5.1",
    )

    go_repository(
        name = "com_github_hashicorp_golang_lru_v2",
        importpath = "github.com/hashicorp/golang-lru/v2",
        sum = "h1:a+bsQ5rvGLjzHuww6tVxozPZFVghXaHOwFs4luLUK2k=",
        version = "v2.0.7",
    )
    go_repository(
        name = "com_github_iancoleman_strcase",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:nTXanmYxhfFAMjZL34Ov6gkzEsSJZ5DbhxWjvSASxEI=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:KwWnWVWCNtNq/ewIX7HIKnELmEx2nDP42yskD/pi7QE=",
        version = "v0.0.0-20240312041847-bd984b5ce465",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:bVoTr12EGANZz66nZPkMInAV/KHD2TxH9npjXXgiB3w=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_github_jackc_pgerrcode",
        importpath = "github.com/jackc/pgerrcode",
        sum = "h1:s+4MhCQ6YrzisK6hFJUX53drDT4UsSW3DEhKn0ifuHw=",
        version = "v0.0.0-20220416144525-469b46aa5efa",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:1HLSx5H+tXR9pW3in3zaztoEwQYRC9SQaYUHjTSUOag=",
        version = "v2.3.3",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:iCEnooe7UlwOQYpKFhBabPMi4aNAfoODPEFNiAnClxo=",
        version = "v0.0.0-20240606120523-5a60cdf6a761",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:y+xUdabmyMkJLyApYuPj38mW+aAIqCe5uuBB51rH3Vw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:xVpYkNR5pk5bMCZGfClbO962UIqVABcAGt7ha1s/FeU=",
        version = "v4.18.2",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:mLoDLV6sonKlvjIEsV56SkWNCnuNv531l94GaIzO+XI=",
        version = "v5.7.2",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:PR8nw+E/1w0GLuRFSmiioY6UooMp6KJv0/61nB7icHo=",
        version = "v2.2.2",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:QH0l3hzAU1tfT3rZCnW5zXl+orbkNMMRGJfdJjHVETg=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_jcmturner_goidentity_v6",
        importpath = "github.com/jcmturner/goidentity/v6",
        sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:x1Sv4HaTpepFkXbt2IkL29DXRf8sOfZXo8eRKh687T8=",
        version = "v8.4.4",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_josharian_intern",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_k0kubun_pp",
        importpath = "github.com/k0kubun/pp",
        sum = "h1:EKhKbi34VQDWJtq+zpsKSEhkHHs9w2P8Izbq8IhLVSo=",
        version = "v2.3.0+incompatible",
    )
    go_repository(
        name = "com_github_kardianos_osext",
        importpath = "github.com/kardianos/osext",
        sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
        version = "v0.0.0-20190222173326-2bc1f35cddc0",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_asmfmt",
        importpath = "github.com/klauspost/asmfmt",
        sum = "h1:4Ri7ox3EwapiOjCki+hw14RyKk201CN4rzyCJRFLpK4=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:c/Cqfb0r+Yi+JtIEq73FWXVkRonBlf0CRNYc8Zttxdo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:tBs3QSyvjDyFTq3uoc/9xFpCuOsJQFNPiAhYdw2skhE=",
        version = "v2.2.10",
    )
    go_repository(
        name = "com_github_knetic_govaluate",
        importpath = "github.com/Knetic/govaluate",
        sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
        version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
    )
    go_repository(
        name = "com_github_knz_go_libedit",
        importpath = "github.com/knz/go-libedit",
        sum = "h1:0pHpWtx9vcvC0xGZqEQlQdfSQs7WRlAjuPvk3fOZDCo=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ktrysmt_go_bitbucket",
        importpath = "github.com/ktrysmt/go-bitbucket",
        sum = "h1:C8dUGp0qkwncKtAnozHCbbqhptefzEd1I0sfnuy9rYQ=",
        version = "v0.6.4",
    )
    go_repository(
        name = "com_github_kujilabo_cocotola_1_23_lib",
        importpath = "github.com/kujilabo/cocotola-1.23/lib",
        sum = "h1:Tmv1dzk0y3HIXNzhEnld4831zfqrTheXpy9X4RcKzdE=",
        version = "v0.0.0-20241124143107-a1ef5177cc84",
    )
    go_repository(
        name = "com_github_kylebanks_depth",
        importpath = "github.com/KyleBanks/depth",
        sum = "h1:5h8fQADFrWtarTdtDudMmGsC7GPbOAu6RVB3ffsVFHc=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:WT9HwE9SGECu3lg4d/dIA+jxlljEa1/ffXKmRjqdmIQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=",
        version = "v1.10.9",
    )
    go_repository(
        name = "com_github_lufia_plan9stats",
        importpath = "github.com/lufia/plan9stats",
        sum = "h1:6E+4a0GO5zZEnZ81pIr0yLvtUWk2if982qA3F3QD6H4=",
        version = "v0.0.0-20211012122336-39d0f177ccd0",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star_v2",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        sum = "h1:sIXJOMrYnQZJu7OB7ANSF4MYri2fTEGIsRLz6LwI4xE=",
        version = "v2.0.4-0.20230330145011-496ad1ac90a4",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_markbates_pkger",
        importpath = "github.com/markbates/pkger",
        sum = "h1:3MPelV53RnGSW07izx5xGxl4e/sdRD6zqseIk0rMASY=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:6Su7aK7lXmJ/U79bYtBjLNaha4Fs1Rg9plHpcH+vvnE=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=",
        version = "v0.0.20",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:2gZY6PC6kBnID23Tichd1K+Z0oS6nE/XwU+Vz/5o4kU=",
        version = "v1.14.22",
    )

    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions_v2",
        importpath = "github.com/matttproud/golang_protobuf_extensions/v2",
        sum = "h1:jWpvCLoY8Z/e3VKvlsiIGKtc+UG6U5vzxaoagmhXfyg=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_micahparks_keyfunc",
        importpath = "github.com/MicahParks/keyfunc",
        sum = "h1:lhKd5xrFHLNOWrDc4Tyb/Q1AJ4LCzQ48GVJyVIID3+o=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_microsoft_go_mssqldb",
        importpath = "github.com/microsoft/go-mssqldb",
        sum = "h1:7cyZ/AT7ycDsEoWPIXibd+aVKFtteUNhDGf3aobP+tw=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:F2VQgta7ecxGYO8k3ZZz3RS8fVIXVxONVUPlNERoyfY=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_minio_asm2plan9s",
        importpath = "github.com/minio/asm2plan9s",
        sum = "h1:AMFGa4R4MiIpspGNG7Z948v4n35fFGB3RR3G/ry4FWs=",
        version = "v0.0.0-20200509001527-cdd76441f9d8",
    )
    go_repository(
        name = "com_github_minio_c2goasm",
        importpath = "github.com/minio/c2goasm",
        sum = "h1:+n/aFZefKZp7spd8DFdX7uMikMLXX4oubIzJF4kv/wI=",
        version = "v0.0.0-20190812172519-36a3d3bbc4f3",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_moby_docker_image_spec",
        importpath = "github.com/moby/docker-image-spec",
        sum = "h1:jMKff3w6PgbfSa69GfNg+zN/XLhfXJGnEx3Nl2EsFP0=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_moby_patternmatcher",
        importpath = "github.com/moby/patternmatcher",
        sum = "h1:GmP9lR19aU5GqSSFko+5pRqHi+Ohk1O69aFiKkVGiPk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_moby_sys_sequential",
        importpath = "github.com/moby/sys/sequential",
        sum = "h1:OPvI35Lzn9K04PBbCLW0g4LcFAJgHsvXsRyewg5lXtc=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_moby_sys_user",
        importpath = "github.com/moby/sys/user",
        sum = "h1:WmZ93f5Ux6het5iituh9x2zAG7NFY9Aqi49jjE1PaQg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_moby_sys_userns",
        importpath = "github.com/moby/sys/userns",
        sum = "h1:tVLXkFOxVu9A64/yh59slHVv9ahO9UIev4JZusOLG/g=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        sum = "h1:xt8Q1nalod/v7BqbG21f8mQPqH+xAaC9C3N3wfWbVP0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_modocache_gover",
        importpath = "github.com/modocache/gover",
        sum = "h1:8Q0qkMVC/MmWkpIdlvZgcv2o2jrlF6zqVOh7W5YHdMA=",
        version = "v0.0.0-20171022184752-b58185e213c5",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:r3y12KyNxj/Sb/iOE46ws+3mS1+MZca1wlHQFPsY/JU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mtibben_percent",
        importpath = "github.com/mtibben/percent",
        sum = "h1:5gssi8Nqo8QU/r2pynCm+hBQHpkB/uNK7BJCFogWdzs=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_mutecomm_go_sqlcipher_v4",
        importpath = "github.com/mutecomm/go-sqlcipher/v4",
        sum = "h1:sV1tWCWGAVlPhNGT95Q+z/txFxuhAYWwHD1afF5bMZg=",
        version = "v4.4.0",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_nakagami_firebirdsql",
        importpath = "github.com/nakagami/firebirdsql",
        sum = "h1:P48LjvUQpTReR3TQRbxSeSBsMXzfK0uol7eRcr7VBYQ=",
        version = "v0.0.0-20190310045651-3c02a58cfed8",
    )
    go_repository(
        name = "com_github_ncruces_go_strftime",
        importpath = "github.com/ncruces/go-strftime",
        sum = "h1:bY0MQC28UADQmHmaF5dgpLmImcShSi2kHU9XLdhx/f4=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_neo4j_neo4j_go_driver",
        importpath = "github.com/neo4j/neo4j-go-driver",
        sum = "h1:fhFP5RliM2HW/8XdcO5QngSfFli9GcRIpMXvypTQt6E=",
        version = "v1.8.1-0.20200803113522-b626aa943eba",
    )
    go_repository(
        name = "com_github_ohler55_ojg",
        importpath = "github.com/ohler55/ojg",
        sum = "h1:sDwc4u4zex65Uz5Nm7O1QwDKTT+YRcpeZQTy1pffRkw=",
        version = "v1.25.0",
    )
    go_repository(
        name = "com_github_oklog_ulid_v2",
        importpath = "github.com/oklog/ulid/v2",
        sum = "h1:+9lhoxAP56we25tyYETBBY1YLA2SaoLvUFgrP2miPJU=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_onrik_gorm_logrus",
        importpath = "github.com/onrik/gorm-logrus",
        sum = "h1:JKeFH+j8AIpCDtsxHgteMtQeZtJ1k+M6UlUXwfkd2+o=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_onrik_logrus",
        importpath = "github.com/onrik/logrus",
        sum = "h1:nWY19z/YktcI9SaZvMocXl7LMtbjIB5RRzKG7JSKIYQ=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:29JGrr5oVBm5ulCWet69zQkzWipVXIol6ygQUe/EzNc=",
        version = "v1.16.4",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:WjP/FQ/sk43MRmnEcT+MlDw2TFvkrXlprrPST/IudjU=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:8SG7/vwALn54lVB/0yZ/MMwhFrPYtpEHQb2IpWsCzug=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_orandin_slog_gorm",
        importpath = "github.com/orandin/slog-gorm",
        sum = "h1:FgA8hJufF9/jeNSYoEXmHPPBwET2gwlF3B85JdpsTUU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_patrickmn_go_cache",
        importpath = "github.com/patrickmn/go-cache",
        sum = "h1:HRMgzkcYKYpi3C8ajMPV8OFXaaRUnok+kx1WdO15EQc=",
        version = "v2.1.0+incompatible",
    )

    go_repository(
        name = "com_github_pborman_getopt",
        importpath = "github.com/pborman/getopt",
        sum = "h1:BHT1/DKsYDGkUgQ2jmMaozVcdk+sVfz0+1ZJq4zkWgw=",
        version = "v0.0.0-20170112200414-7148bc3a4c30",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:YmeHyLY8mFWbdkNWwpr+qIL2bEqT0o95WSdkNHvL12M=",
        version = "v2.2.3",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:xaKrnTkyoqfh1YItXl56+6KJNVYWlEEPuAQW9xsplYQ=",
        version = "v4.1.18",
    )
    go_repository(
        name = "com_github_pkg_browser",
        importpath = "github.com/pkg/browser",
        sum = "h1:+mdjkGKdHQG3305AYmdv1U2eRNDiU2ErMBj1gwrq8eQ=",
        version = "v0.0.0-20240102092130-5ac0b6a4141c",
    )
    go_repository(
        name = "com_github_pkg_diff",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        importpath = "github.com/pkg/profile",
        sum = "h1:hnbDkaNWPCLMO9wGLdBFTIZvzDrDfBM2072E1S9gJkA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:1sLMdKq4gNANTj0dUibycTLzpIEKVnLnbaEkxws78nw=",
        version = "v0.6.1-0.20241121165744-79df5c4772f2",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_power_devops_perfstat",
        importpath = "github.com/power-devops/perfstat",
        sum = "h1:ncq/mPwQF4JjgDlrVEn3C11VoGHZN7m8qihwgMEtzYw=",
        version = "v0.0.0-20210106213030-5aafc221ea8c",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:DIsaGmiaBkSangBgMtWdNfxbMNdku5IK6iNhrEqWvdA=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:ZKSh/rekM+n3CeS952MLRAdFwIKqeY8b62p8ais2e9E=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:xasJaQlnWAeyHdUBeGjXmutelfJHWMRr+Fg4QszZ2Io=",
        version = "v0.62.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:YagwOFzUgYfKKHX6Dr+sHT7km/hxC76UB0learggepc=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_remyoudompheng_bigfft",
        importpath = "github.com/remyoudompheng/bigfft",
        sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
        version = "v0.0.0-20230129092748-24d4a6f8daec",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:KvO1DLK/DRN07sQ1LQKScxyZJuNnedQ5/wKSR38lUII=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_rqlite_gorqlite",
        importpath = "github.com/rqlite/gorqlite",
        sum = "h1:V7x0hCAgL8lNGezuex1RW1sh7VXXCqfw8nXZti66iFg=",
        version = "v0.0.0-20230708021416-2acd02b70b79",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:KqfZb0pUVN2lYqZUYRddxF4OR8ZMURnJIG5Y3VRLtww=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_samber_slog_gin",
        importpath = "github.com/samber/slog-gin",
        sum = "h1:6DMAcy2gBFyyztrpYIvAcXZH1sA/j75iSSXuqhirLtg=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_shirou_gopsutil_v3",
        importpath = "github.com/shirou/gopsutil/v3",
        sum = "h1:z90NtUkp3bMtmICZKpC4+WaknU1eXtp5vtbQ11DgpE4=",
        version = "v3.23.12",
    )
    go_repository(
        name = "com_github_shoenig_go_m1cpu",
        importpath = "github.com/shoenig/go-m1cpu",
        sum = "h1:nxdKQNcEB6vzgA2E2bvzKIYRuNj7XNJ4S/aRSwKzFtM=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_shoenig_test",
        importpath = "github.com/shoenig/test",
        sum = "h1:kVTaSd7WLz5WZ2IaoM0RSzRsUD+m8wRR+5qvntpn4LU=",
        version = "v0.6.4",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_snowflakedb_gosnowflake",
        importpath = "github.com/snowflakedb/gosnowflake",
        sum = "h1:KSHXrQ5o7uso25hNIzi/RObXtnSGkFgie91X82KcvMY=",
        version = "v1.6.19",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:EaGW2JJh15aKOejeuJ+wpFSHnbd7GE6Wvp3TsNhb6LY=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:Xv5erBjTwe/5IxqUQTdXv5kgmIvbHo3QQyRwhJsOfJA=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_swaggo_files",
        importpath = "github.com/swaggo/files",
        sum = "h1:1gGXVIeUFCS/dta17rnP0iOpr6CXFwKD7EO5ID233e4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_swaggo_gin_swagger",
        importpath = "github.com/swaggo/gin-swagger",
        sum = "h1:8mWmHLolIbrhJJTflsaFoZzRBYVmEE7JZGIq08EiC0Q=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_swaggo_swag",
        importpath = "github.com/swaggo/swag",
        sum = "h1:kHtaBe/Ob9AZzAANfcn5c6RyCke9gG9QpH0jky0I/sA=",
        version = "v1.8.9",
    )
    go_repository(
        name = "com_github_tcolgate_mp3",
        importpath = "github.com/tcolgate/mp3",
        sum = "h1:XQdibLKagjdevRB6vAjVY4qbSr8rQ610YzTkWcxzxSI=",
        version = "v0.0.0-20170426193717-e79c5a46d300",
    )
    go_repository(
        name = "com_github_testcontainers_testcontainers_go",
        importpath = "github.com/testcontainers/testcontainers-go",
        sum = "h1:5fbgF0vIN5u+nD3IWabQwRybuB4GY8G2HHgCkbMzMHo=",
        version = "v0.34.0",
    )
    go_repository(
        name = "com_github_testcontainers_testcontainers_go_modules_mysql",
        importpath = "github.com/testcontainers/testcontainers-go/modules/mysql",
        sum = "h1:Tqz17mGXjPORHFS/oBUGdeJyIsZXLsVVHRhaBqhewGI=",
        version = "v0.34.0",
    )

    go_repository(
        name = "com_github_tidwall_gjson",
        importpath = "github.com/tidwall/gjson",
        sum = "h1:uo0p8EbA09J7RQaflQ1aBRffTR7xedD2bcIVSYxLnkM=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_github_tidwall_match",
        importpath = "github.com/tidwall/match",
        sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:RWIZEg2iJ8/g6fDDYzMpobmaoGh5OLl4AXtGUGPcqCs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_tklauser_go_sysconf",
        importpath = "github.com/tklauser/go-sysconf",
        sum = "h1:0QaGUFOdQaIVdPgfITYzaTegZvdCjmYO52cSFAEVmqU=",
        version = "v0.3.12",
    )
    go_repository(
        name = "com_github_tklauser_numcpus",
        importpath = "github.com/tklauser/numcpus",
        sum = "h1:ng9scYS7az0Bk4OZLvrNXNSAO2Pxr1XXRAPyjhIx+Fk=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_twitchyliquid64_golang_asm",
        importpath = "github.com/twitchyliquid64/golang-asm",
        sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:9LC83zGrHhuUA9l16C9AHXAqEV/2wBQ4nkvumAE65EE=",
        version = "v1.2.12",
    )
    go_repository(
        name = "com_github_xanzy_go_gitlab",
        importpath = "github.com/xanzy/go-gitlab",
        sum = "h1:rWtwKTgEnXyNUGrOArN7yyc3THRkpYcKXIXia9abywQ=",
        version = "v0.15.0",
    )
    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:VOMT+81stJgXW3CpHyqHN3AXDYIMsx56mEFrB37Mb/E=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:kdwGpVNwPFtjs98xCGkHjQtGKh86rDcRZN17QEMCOIs=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
        version = "v0.0.0-20180127040702-4e3ac2762d5f",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
        version = "v0.0.0-20181117223130-1be2e3e5546d",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_github_yusufpapurcu_wmi",
        importpath = "github.com/yusufpapurcu/wmi",
        sum = "h1:E1ctvB7uKFMOJw3fdOW32DwGE9I7t++CRUEMKvFoFiw=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_zeebo_xxh3",
        importpath = "github.com/zeebo/xxh3",
        sum = "h1:xZmwmqxHZA8AI603jOQ0tMqmBr9lPeFwGg6d+xy9DC0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_gitlab_nyarla_go_crypt",
        importpath = "gitlab.com/nyarla/go-crypt",
        sum = "h1:7gd+rd8P3bqcn/96gOZa3F5dpJr/vEiDQYlNb/y2uNs=",
        version = "v0.0.0-20160106005555-d9a5dc2b789b",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:jsypSnrE/w4mJysioGdMBg4MiW/hHx/sArFpaBWHdME=",
        version = "v0.118.3",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:axlU03FRiXDNupsmPG7LKzuS4Enk1gf598M62lWVB74=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:8zVoeiBa4erMCLEXltOcqVEsZhS26JZ5/Vrgs59eQiI=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:IuxpcwpkmexB7zvzwkCDsKok8mwPWOn2kFaHQYi1INA=",
        version = "v1.73.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:O2kWr2Sd4ep3I+YJ4aiY0G4+zWz6sp4eTce+JVns9TM=",
        version = "v0.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:Mn7cC5iWJz+cSMS/Hb+N2410CpZ6c8XpJKaexBl0Gxs=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:Wlr+30Tha0SMCvQYZKdrh+HkpOyl0CQFSlzeY/Gg1gs=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:j9CJg/oC884OX5cDpiwNt1ZlDXNV6Zb9Mp1YmRrOG0k=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:jrcanSzj9J1erevZuxldvsDwY+0k/DeFFzlnSfPGfL8=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:dPQ07rW4eku8OgNWDOaQaVGcE4+XfhH8BSbVwdVQ+wU=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:ZNXGB6+T7VmWdf6//VqxLdZ/sk0no8W0ujanHeJwDRw=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:6oNgjcs5KCPGBD71G0IccK6TfeFsEtBTyQ3Q+Dn09bs=",
        version = "v1.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:RU1WhF1zMggdXAZ+ezYTn4Eh/FdiX7sz8lLXGERn4Po=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_auth",
        importpath = "cloud.google.com/go/auth",
        sum = "h1:Ly0u4aA5vG/fsSsxu98qCQBemXtAtJf+95z9HK+cxps=",
        version = "v0.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_auth_oauth2adapt",
        importpath = "cloud.google.com/go/auth/oauth2adapt",
        sum = "h1:/Lc7xODdqcEw8IrZ9SvwnlLX6j9FHQM74z6cBk9Rw6M=",
        version = "v0.2.7",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:vkD+hQ75SMINMgJBT/KDpFYvfQLzJbtIQZdw0AWq8Rs=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:OL+KT+wCumdDhG44aeqGAdkwdT8Wa4Lh+o4INM+CQjw=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:lXuTaELvU0P0ARbTFxxdpOC/dFnZZeGglSw06BtO//8=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:ezavJc0Gzh4N8zBskO/DnUVMWPa8lqH/tmQSyaknmCA=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:EKOSqjtO7jPpJoEzDmRctGea3c2EOGoexy8VyY9dNro=",
        version = "v1.66.2",
    )
    go_repository(
        name = "com_google_cloud_go_bigtable",
        importpath = "cloud.google.com/go/bigtable",
        sum = "h1:UEacPwaejN2mNbz67i1Iy3G812rxtgcs6ePj1TAg7dw=",
        version = "v1.35.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:xMlO3hc5BI0s23tRB40bL40xSpxUR1x3E07Y5/VWcjU=",
        version = "v1.20.1",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:X8JRfmk0/vyRqLusEyAPr0nZCK6RKae9omB4lrit0XI=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:2UP31fg7b+y3F0OmNbPHOKPEJ+6LOMfxAXX4p8xGCy4=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:oHyO3QAZ6kdf6SwqnUTBz50ND6Nk2rxZtboUiF4dgLE=",
        version = "v1.19.2",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:zmDznviZpvkCla0adbp7jJsMYZ9bABCbcPK2cBUHwg8=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:T/rkkKE0KhQFMcO3+QWL82xakA9kRumLXY1lq5adIts=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:rXdznKjCa7WpzmvR2plrn2KJ+RZC1oYxPiRWNQjjf3k=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:+k/kmViu4TEi97NGaxAATYtpYBviOWJySPZ+ekA95kk=",
        version = "v1.34.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:A6hENjEsCDtC1k8byVsgwvVcioamEHvZ4j01OwKxG9I=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:xJoZbX0HM1zht8KxAB38hs2v4Hcl+vXGLo454LrdwxA=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:8ncSEBjkng6ucCICauaUGzBomoM2VyYzleAum1OFcow=",
        version = "v1.42.2",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:1D8U75BeotZxrG4jR6NYBtOt+uAeBsWhpBZmSYLakQw=",
        version = "v0.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:3bAfstDB6rlHyK0TvqxEwaeOvoN9UgCs2bn03+VXmss=",
        version = "v1.24.3",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:+7IfIXzYWSybIIDGK9FN2uqBsP/5b/Y0pBYzNhcmKSU=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:ZpGkZV8OyhUhvN/tfLffU2ki5ERTtqOunkIaiVAhmw0=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:FTMtsf2nfGGlDCuE84/RvVaCcTIYE7WQSB0noeO0cwI=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:PqoA3gnOWaLcHCnqoZe4jh3jmiv6+Z7W2xUUkw/j4jE=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:j4hD6opb+gq9CJNPFIlIggoW8Kjymg8Wmy2mdHmQoiw=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:2vOv471LrcSn91VNzijcH+OkDRLa3kdyymOfKqbwZ4c=",
        version = "v2.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:lGUj2FYs650EUPDMV6plWBAoh8qH9Bu1KCz1PUYF2VY=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:NNpXoyEqIJmZFc0ACcwBEaXnmscUpcG4NkKnbCePmiM=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:C5AeEdze55feJVb17a40QmlnyH/aMhn/uf3Go3hIqPA=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:1c2Cd3jdb0mrKHHfyzSQ5DRmxgYd07tIZZzuMNrwDxU=",
        version = "v1.26.2",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:JIl3noWZKV16qb1tXTWyUH1IHx9c0lQUCo4xFMTYjIA=",
        version = "v1.65.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:9kz7+gaB/0gBZsDUnNT1asDihNZSrRFSeUTBcBdUAkk=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:hswVobCWUTXtmn+4QqUIVkai7sDOe0QS2KB3IpqLkik=",
        version = "v1.35.2",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:wnqN5YwMrtLSjn+HB2sChgmZ6iocOta4Q41giQsiRjY=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:SwQuHQiheVfL7b5ar/AXDberiaqr/yiue8X55AdWnZU=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:isaoPwWX8kbAOea4qahcmttoS79+gQhvKsfg5L5AgH8=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:Paw495vxVyKuAgcQ2NQk09iRZBhPYRytknydEnvzcv4=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:RMymT7R87LaxKugOKwooOoheWXUm1NMeOfh3CVU9g54=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:vTXQI5qYKZ8dmCyHN+zVfaMyXCYbyZNM0CkPzpPUn7Q=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:cuydCaLS7Vl2SatAeivXyhbhDEIR8BDmtn4egDhIn2s=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:V0vCHSgFTUqKn57+PUXp1UfQY0/aMkveAw7wXeM3Lq0=",
        version = "v1.19.3",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:djdExe/QgoKdp1gnIO1G5BoO1o/yGQOQJJEZ4QKTEXQ=",
        version = "v1.6.3",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:YVpR0vlHSP/wD74PXEbKua4Aamud+wiYm4TiewNjD3M=",
        version = "v0.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:yZ6lNJ9rNIoQmWrG14dB3+BFjS/EIRBf7Bo6jc5QWlE=",
        version = "v0.15.3",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:JWe6PDNpNU88ZYvQkTd7w28fgeIs/gg6i0hcjUkgZ3M=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_google_cloud_go_grafeas",
        importpath = "cloud.google.com/go/grafeas",
        sum = "h1:CobnwnyeY1j1Defi5vbEircI+jfrk3ci5m004ZjiFP4=",
        version = "v0.3.11",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:f3eMYsCDdg2AeldIPdKmBRxN1WoiTpE3RvX5orcm/I8=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:ZNfy/TYfn2uh/ukvhp783WhnbVluqf/tzOaqVUPlIPA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:OWNYFHPyIBNHEAEFdVKOltYWe0g3izSrpFJW6Iidovk=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:wbFF7twu0XScFr+dtsVxTTttbFIRYt/SJjZiHFidtYE=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:aPWYQ+A1NX6ou/5U0nFAiXWdVT8OBxZYVZt2fBl2gWA=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:x3EeWKuYwdlo2HLse/876ZrKjk2L5r7Uexfm8+p6mSI=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:8hmFMiS3wjjj3TX/U1zZYTgzwZoUjDbo9PaqcYEmuB4=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:Z05C+Ui953f0EQx9hJ1la6+QQl8ADrIs3iNwP5Elkpg=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:7j0HgAp0B94o1YRDqiqm26w4q1rDMH7XNRU34lJXHYc=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:3tyw9rO3E2XVXzSApn1gyEEnH2K9SynNQjMlBi3uHLg=",
        version = "v0.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:b9xGs24BIjfyvLgCtJoClOZpPi8d8owPgWe5JEINgaY=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:Cw4zvH6YLWiNpmpgtTteP7uopo0SSGDZgPU1n/whBLo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:nRBjeaMLipw05Br+qDAlSCcCQAAlat4mvpafztbEVgc=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:XH/qT3GbbSH//R0JTqR77lRpBxaa0N9sHgAzfwbTrv0=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:jDqeCw6NGDRAPT9+2Y/EjnWAB0BfCcUfmPLOyhB0eHs=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:csSKiCJ+WVRgNkRzzz3BPoGjFhjPY23ZTcaenToJxMM=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:YsVhG71ZC4FkqCP2oCI55x/JeGFyd7738Lt8iNTrzJw=",
        version = "v1.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:oEoFGPYxTBsY47h0zdoE2ojV5aU/541D83UmxfjHWaE=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:JLJBFbxc8D7/OS81MyRoKhc2OvnVJxy5VMoQqqAhA7k=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:+9DrGJcZhCu6B2t0JJorekjIUBvg/KvBmXJYGmfvVvA=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:JwQjjoBZJpsoMQe/3mhVBMVZuSdagHg2pGOnwh2Jk+E=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:SFAsKyqvtS8VFcsq+JgXAeRkrksB9UH+AH7iFamkmlc=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:WFvgmjq/FO5GiXlhebltA9N14KdbLMcgG88ME+SWeBo=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:cyf1PMK5c2/WOIr5r2lxjH/XBJMA9P4zC8Tm10i0z3M=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:yomxnFPk+ye0zd0mJ15nn9fH4Ns7ex4xA3ll+u2q59A=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:T5mGFV0ggBKg3qt9myFRiGJu+nIUucuHLAtVpAuQ08I=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:ekIWI8JbKkpOfrgH/THGamQE/D16tcVBYJyrkseVcYI=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:fu2LABMi7CgZORQ2oNGbc0hoZ0FTqLkjGqIgAV/Kc7U=",
        version = "v0.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:Ou2Qu4INnf7ykrFjGv2ntFOjVo8Nloh/+OffF4mUu9w=",
        version = "v1.47.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:jLQozsEVr+c6tOU13vDugtnaBSUy/PD5zK6mhm+uF1Y=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:T5YGzaXwTesHaPDNTAuU3neDwZEnfjce70zufPFUwno=",
        version = "v2.19.4",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:kBpcYPx4ys4lrDGKp4OhP2uy8h7UjlmLW/qoO5Xb2bY=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:dVlOjxsbjuhlwu4MIcyPWe09qVcDqc419iOjdPl5RHk=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:xcu35SCyHSp+nKV6QNIklgkBKTH1qb0aLUXjl0mSR8I=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:SHOMw0kX0xWratC5Vb5VULBeWiGlPYAs82kiZqNtWpM=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:13HOFU7v4cEvIHXSAQbinF4wp2Baybbq7q9FMctg1Ek=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:PT6CUlazIFIOLLJnV+bPBtiSH8iusKZ+FZRzZYFt2vk=",
        version = "v1.19.2",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:9WeTqeEcriXqRViXMNwczjFJjixOSBlSlk/fW3lfKPg=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:ewVvigBnEnrr9Ih8CKnLVoB5IiULaWfYU5nEnnfVAto=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:W++V0EL9iL6T2+ec24Dm++bIti0tI6Gx6sCosDBters=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:ya9gfY1ign6Yy25VMMMgZ9xy7D/TczDB0ElXcyWmEVE=",
        version = "v1.18.3",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:IdDiAa7gYtL7Gdx+wEaNHimudk3ZkEGNhdz9FuEuxWM=",
        version = "v1.36.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:oFkCp6ti7fc7hzeROmOPQuPBHFqwyhcsv3Yrma28+uc=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:mjYgUsOtV3jl9xvDmcvlRRmA64deEPf52zOfuc68b/g=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:vYbVZuXfnFwvNcvH3lhI2PeUA+kHyqKmLC7mJWaC4Ok=",
        version = "v1.76.1",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:qvURtJs7BQzQhbxWxwai0pT79S8KLVKJ/4W8igVkt1Y=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:3TbVkzTooBvnZsk7WaAQfOsNrdoM8QHusXA1cpk6QJs=",
        version = "v1.50.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:W3v9A7MGBN7H9sAFstyciwP/1XEQhUhZfrjclmDnpMs=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:olv+s2g+LGXeJi+MYF1wI44/TwHaVnO0N7PiucVf5ZQ=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:YF/RdNb+jUEp22cIZCvqiFjfA5OxGE+Dxss3mhXU7oQ=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:BvMNijOb6Vd46Rr/SR5jWv1MPosOhVsi0UaeAGNjeds=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:c+I4YFjxRQjvAhRmSsmjpASUKq88chOX854ied0K/pE=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:XJ7LipYJi80BCgVk2lx1fwc7DIYM6oV2qx1G4IAGQ5w=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:C2FH+6yr6LCZC4fP0gm9FwJB/SRh5Ul88O5Sc/bL83I=",
        version = "v1.23.3",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:zNTOUQyatGQtnCJ2dR3faRtpWQOlC8wszJqwG5CtwVM=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:dPvfDuPqPH+Yscf0f2f1RprvKkoo+N/j0a+IbLYX7Cs=",
        version = "v2.9.3",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:dpCQq3pj2HnKdbvGTftdWymm3r4ovF7JW5z8xBcO2x4=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:TfuQr5j7qriINulUMotaC/+27SQaW2thIkF3Gb6VJ38=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:vxVaoFM64M/ht619c4wZNF0iq0QPaMWElOh7Ns4r41A=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:yh0v/5n49VO4/i9pYfDm1gLJUj1Ph3Xzegn8WvK9YRA=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:/uxhVCWKXzPw5pVfnBOVjaSiQ6Bm0tDExDOCLV40thw=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:lNFDMranJymDEB7cTI7DI9czbc1WU0RWY9KCEv9zuDY=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_google_firebase_go_v4",
        importpath = "firebase.google.com/go/v4",
        sum = "h1:KJtV4rAfO2CVCp40hBfVk+mqUqg7+jQKx7yOgFDnXBg=",
        version = "v4.15.2",
    )
    go_repository(
        name = "com_lukechampine_uint128",
        importpath = "lukechampine.com/uint128",
        sum = "h1:cDdUVfRwDUDovz610ABgFD17nXD4/uDgVHl2sC3+sbo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_nullprogram_x_optparse",
        importpath = "nullprogram.com/x/optparse",
        sum = "h1:xGFgVi5ZaWOnYdac2foDT3vg0ZZC9ErXFV57mr4OHrI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )

    go_repository(
        name = "com_zombiezen_go_sqlite",
        importpath = "zombiezen.com/go/sqlite",
        sum = "h1:N1s3RIljwtp4541Y8rM880qgGIgq3fTD2yks1xftnKU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "dev_cel_expr",
        importpath = "cel.dev/expr",
        sum = "h1:o+Wj235dy4gFYlYin3JsMpp3EEfMrPm/6tdoyjT98S0=",
        version = "v0.21.2",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_crawshaw_iox",
        importpath = "crawshaw.io/iox",
        sum = "h1:yDf7ARQc637HoxDho7xjqdvO5ZA2Yb+xzv/fOnnvZzw=",
        version = "v0.0.0-20181124134642-c51c3df30797",
    )
    go_repository(
        name = "io_filippo_edwards25519",
        importpath = "filippo.io/edwards25519",
        sum = "h1:FNf4tywRC1HmFuKW5xopWpigGjJKiJSV0Cqo0cJWDaA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_gorm_driver_mysql",
        importpath = "gorm.io/driver/mysql",
        sum = "h1:MndhOPYOfEp2rHKgkZIhJ16eVUIRf2HmzgoPmh7FCWo=",
        version = "v1.5.7",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:ubBVAfbKEUld/twyKZ0IYn9rSQh448EdelLYk9Mv314=",
        version = "v1.5.11",
    )
    go_repository(
        name = "io_gorm_driver_sqlserver",
        importpath = "gorm.io/driver/sqlserver",
        sum = "h1:xA+Y1KDNspv79q43bPyjDMUgHoYHLhXYmdFcYPobg8g=",
        version = "v1.5.4",
    )
    go_repository(
        name = "io_gorm_gorm",
        importpath = "gorm.io/gorm",
        sum = "h1:I0u8i2hWQItBq1WfE0o2+WuL9+8L21K9e2HHSTE/0f8=",
        version = "v1.25.12",
    )
    go_repository(
        name = "io_gorm_plugin_dbresolver",
        importpath = "gorm.io/plugin/dbresolver",
        sum = "h1:wFwINGZZmttuu9h7XpvbDHd8Lf9bb8GNzp/NpAMV2wU=",
        version = "v1.5.3",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_auto_sdk",
        importpath = "go.opentelemetry.io/auto/sdk",
        sum = "h1:cH53jehLUN6UFLY71z+NDOiNJqDdPRaXzTel0sJySYA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_detectors_gcp",
        importpath = "go.opentelemetry.io/contrib/detectors/gcp",
        sum = "h1:JRxssobiPg23otYU5SbWtQC//snGVIM3Tx6QRzlQBao=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_github_com_gin_gonic_gin_otelgin",
        importpath = "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin",
        sum = "h1:5Acs0t57/EJbB54SUEdALa+0ln2UEawYPUSIX3qdE14=",
        version = "v0.59.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:rgMkmiGfix9vFJDcDi1PK8WEQP4FLQwLDfhp5ZLpFeE=",
        version = "v0.59.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:CV7UdSGJt/Ao6Gp4CXckLxVRRsRgDHoI8XjbL3PDl8s=",
        version = "v0.59.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_propagators_b3",
        importpath = "go.opentelemetry.io/contrib/propagators/b3",
        sum = "h1:MazJBz2Zf6HTN/nK/s3Ru1qme+VhWU5hm83QxEP+dvw=",
        version = "v1.32.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:zRLXxLCgL1WyKsPVrgbSdMN4c0FMkDAskSTQP+0hdUY=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_jaeger",
        importpath = "go.opentelemetry.io/otel/exporters/jaeger",
        sum = "h1:ES8/j2+aB+3/BUw51ioxa50V9btN1eew/2J7N7n1tsE=",
        version = "v1.11.2",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:OeNbIYk/2C15ckl7glBlOBp5+WlYsOElzTNmiPW/x60=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
        sum = "h1:BEj3SPM81McUZHYjRS5pEgNgnmzGJ5tRpU5krWnV8Bs=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdoutmetric",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric",
        sum = "h1:WDdP9acbMYjbKIyJUhTvtzj601sVJOqgWdUxSdR/Ysc=",
        version = "v1.29.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdouttrace",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdouttrace",
        sum = "h1:jBpDk4HAUsrnVO1FsfCfCOTEc/MkInJmvfCHYLFiT80=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:+eTR3U0MyfWjRDhmFMxe2SsW64QrZ84AOhvqS7Y+PoQ=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:95zS4k/2GOy069d321O8jWgYsW3MzVV+KuSPKp7Wr1A=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:5CeK9ujjbFVL5c1PhLuStg1wxA7vQv7ce1EK0Gyvahk=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:+ouXS2V8Rd4hp4580a8q23bg0azF2nI8cqLYnC8mh/k=",
        version = "v1.34.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:xJvq7gMzB31/d406fB8U5CBdyQGw4P399D1aQWU/3i4=",
        version = "v1.5.0",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_pdf",
        importpath = "rsc.io/pdf",
        sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:JUTaWEriXmEy5AhvdMgksGGPEFsYfUKaPEYXd4c3Wvc=",
        version = "v0.223.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_appengine_v2",
        importpath = "google.golang.org/appengine/v2",
        sum = "h1:LvPZLGuchSBslPBp+LAhihBeGSiRh1myRoYK4NtuBIw=",
        version = "v2.0.6",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:EZ4nXs4XXUdRhv/pmiWlz5Hb2pbbPrHruKIN+v8UY+A=",
        version = "v0.0.0-20250227231956-55c901821b1e",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:nsxey/MfoGzYNduN0NN/+hqP9iiCIYsrVbXb/8hjFM8=",
        version = "v0.0.0-20250227231956-55c901821b1e",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:UZtupsOaDeUm4KiG4HQTSyENUuCayW8K5d5cs7zK79c=",
        version = "v0.0.0-20250219182151-9fdb1cabc7b2",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:YA5lmSs3zc/5w+xsRcHqpETkaYyK63ivEPzNTcUUlSA=",
        version = "v0.0.0-20250227231956-55c901821b1e",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:pWFv03aZoHzlRKHWicjsZytKAiYCtNS0dHbXnIdq7jQ=",
        version = "v1.70.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:tPhr+woSbjfYvY6/GPufUoYizxw1cF/yFoxJ2fmpwlM=",
        version = "v1.36.5",
    )
    go_repository(
        name = "org_golang_x_arch",
        importpath = "golang.org/x/arch",
        sum = "h1:z9JUEZWr8x4rR0OU6c4/4t6E6jOZ8/QBS2bBYBm4tx4=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:b15kiHdrGCHrP6LvwaQ3c03kgNhhiMgvlhxHQhmg2Xs=",
        version = "v0.35.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:aWwlzYV971S4BXRS9AmqwDLAD85ouC6X+pocatKY58c=",
        version = "v0.0.0-20250228200357-dead58393ab7",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
        version = "v0.0.0-20190802002840-cff245a6509b",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:2M3HP5CCK1Si9FQhwnzYhXdG6DXeebvUHFpre8QvbyI=",
        version = "v0.0.0-20201208152925-83fdc39ff7b5",
    )
    go_repository(
        name = "org_golang_x_mobile",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:Zb7khfcRGKk+kqfxFaP5tZqCnDZMjC5VtUBs87Hr6QM=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:T5GQRQb2y08kTAByq9L4/bz8cipCdA8FbRTXewonqY8=",
        version = "v0.35.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:da9Vo7/tDv5RH/7nZDz1eMGS/q1Vv1N/7FCrBhI9I3M=",
        version = "v0.27.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:GGz8+XQP4FvTTrjZPzNKTMFtSXH80RAzG+5ghFPgK9w=",
        version = "v0.11.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:QjkSwP/36a20jFYWkSue1YwXzLmsV5Gfq7Eiy72C1uc=",
        version = "v0.30.0",
    )
    go_repository(
        name = "org_golang_x_telemetry",
        importpath = "golang.org/x/telemetry",
        sum = "h1:zf5N6UOrA487eEFacMePxjXAJctxKmyjKUsjA11Uzuk=",
        version = "v0.0.0-20240521205824-bda55230c457",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:L6pJp37ocefwRRtYPKSWOWzOtWSxVajvz2ldH/xi3iU=",
        version = "v0.29.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:bofq7m3/HAFvbF51jz3Q9wLg3jkvSPuiZu/pD1XwgtM=",
        version = "v0.22.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:3usCWA8tQn0L8+hFJQNgzpWbd89begxN66o1Ojdn5L4=",
        version = "v0.10.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:BgcpHewrV5AUp2G9MebG4XPFI1E2W41zU1SaqVA9vJY=",
        version = "v0.30.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:noIWHXmPHxILtqtCOPIhSt0ABwskkZKjD3bXGnZGpNY=",
        version = "v0.0.0-20240903120638-7835f813f4da",
    )
    go_repository(
        name = "org_modernc_b",
        importpath = "modernc.org/b",
        sum = "h1:vpvqeyp17ddcQWF29Czawql4lDdABCDRbXRAS4+aF2o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_cc_v3",
        importpath = "modernc.org/cc/v3",
        sum = "h1:QoR1Sn3YWlmA1T4vLaKZfawdVtSiGx8H+cEojbC7v1Q=",
        version = "v3.41.0",
    )
    go_repository(
        name = "org_modernc_cc_v4",
        importpath = "modernc.org/cc/v4",
        sum = "h1:TFkx1s6dCkQpd6dKurBNmpo+G8Zl4Sq/ztJ+2+DEsh0=",
        version = "v4.24.4",
    )
    go_repository(
        name = "org_modernc_ccgo_v3",
        importpath = "modernc.org/ccgo/v3",
        sum = "h1:o3OmOqx4/OFnl4Vm3G8Bgmqxnvxnh0nbxeT5p/dWChA=",
        version = "v3.17.0",
    )
    go_repository(
        name = "org_modernc_ccgo_v4",
        importpath = "modernc.org/ccgo/v4",
        sum = "h1:Z2N+kk38b7SfySC1ZkpGLN2vthNJP1+ZzGZIlH7uBxo=",
        version = "v4.23.16",
    )
    go_repository(
        name = "org_modernc_db",
        importpath = "modernc.org/db",
        sum = "h1:2c6NdCfaLnshSvY7OU09cyAY0gYXUZj4lmg5ItHyucg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_file",
        importpath = "modernc.org/file",
        sum = "h1:9/PdvjVxd5+LcWUQIfapAWRGOkDLK90rloa8s/au06A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_fileutil",
        importpath = "modernc.org/fileutil",
        sum = "h1:gQ5SIzK3H9kdfai/5x41oQiKValumqNTDXMvKo62HvE=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_modernc_gc_v2",
        importpath = "modernc.org/gc/v2",
        sum = "h1:aJVhcqAte49LF+mGveZ5KPlsp4tdGdAOT4sipJXADjw=",
        version = "v2.6.3",
    )
    go_repository(
        name = "org_modernc_gc_v3",
        importpath = "modernc.org/gc/v3",
        sum = "h1:d0JExN5U5FjUVHCP6L9DIlLJBZveR6KUM4AvfDUL3+k=",
        version = "v3.0.0-20241223112719-96e2e1e4408d",
    )
    go_repository(
        name = "org_modernc_golex",
        importpath = "modernc.org/golex",
        sum = "h1:wWpDlbK8ejRfSyi0frMyhilD3JBvtcx2AdGDnU+JtsE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_httpfs",
        importpath = "modernc.org/httpfs",
        sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_modernc_internal",
        importpath = "modernc.org/internal",
        sum = "h1:XMDsFDcBDsibbBnHB2xzljZ+B1yrOVLEFkKL2u15Glw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_libc",
        importpath = "modernc.org/libc",
        sum = "h1:3LRd6ZO1ezsFiX1y+bHd1ipyEHIJKvuprv0sLTBwLW8=",
        version = "v1.61.13",
    )
    go_repository(
        name = "org_modernc_lldb",
        importpath = "modernc.org/lldb",
        sum = "h1:6vjDJxQEfhlOLwl4bhpwIz00uyFK4EmSYcbwqwbynsc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_mathutil",
        importpath = "modernc.org/mathutil",
        sum = "h1:GCZVGXdaN8gTqB1Mf/usp1Y/hSqgI2vAGGP4jZMCxOU=",
        version = "v1.7.1",
    )
    go_repository(
        name = "org_modernc_memory",
        importpath = "modernc.org/memory",
        sum = "h1:cL9L4bcoAObu4NkxOlKWBWtNHIsnnACGF/TbqQ6sbcI=",
        version = "v1.8.2",
    )
    go_repository(
        name = "org_modernc_opt",
        importpath = "modernc.org/opt",
        sum = "h1:2kNGMRiUjrp4LcaPuLY2PzUfqM/w9N23quVwhKt5Qm8=",
        version = "v0.1.4",
    )
    go_repository(
        name = "org_modernc_ql",
        importpath = "modernc.org/ql",
        sum = "h1:bIQ/trWNVjQPlinI6jdOQsi195SIturGo3mp5hsDqVU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_sortutil",
        importpath = "modernc.org/sortutil",
        sum = "h1:+xyoGf15mM3NMlPDnFqrteY07klSFxLElE2PVuWIJ7w=",
        version = "v1.2.1",
    )
    go_repository(
        name = "org_modernc_sqlite",
        importpath = "modernc.org/sqlite",
        sum = "h1:EQXNRn4nIS+gfsKeUTymHIz1waxuv5BzU7558dHSfH8=",
        version = "v1.36.0",
    )
    go_repository(
        name = "org_modernc_strutil",
        importpath = "modernc.org/strutil",
        sum = "h1:UneZBkQA+DX2Rp35KcM69cSsNES9ly8mQWD71HKlOA0=",
        version = "v1.2.1",
    )
    go_repository(
        name = "org_modernc_tcl",
        importpath = "modernc.org/tcl",
        sum = "h1:C4ybAYCGJw968e+Me18oW55kD/FexcHbqH2xak1ROSY=",
        version = "v1.15.2",
    )
    go_repository(
        name = "org_modernc_token",
        importpath = "modernc.org/token",
        sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_z",
        importpath = "modernc.org/z",
        sum = "h1:zDJf6iHjrnB+WRD88stbXokugjyc0/pB91ri1gO6LZY=",
        version = "v1.7.3",
    )
    go_repository(
        name = "org_modernc_zappy",
        importpath = "modernc.org/zappy",
        sum = "h1:dPVaP+3ueIUv4guk8PuZ2wiUGcJ1WUVvIheeSSTD0yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:ny3p0reEpgsR2cfA5cjgwFZg3Cv/ofFh/8jbhGtz9VI=",
        version = "v1.7.5",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_goleak",
        importpath = "go.uber.org/goleak",
        sum = "h1:2K3zAYmnTNqV73imy9J1T3WC+gmCePx2hEGkimedGto=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "tech_einride_go_aip",
        importpath = "go.einride.tech/aip",
        sum = "h1:16/AfSxcQISGN5z9C5lM+0mLYXihrHbQ1onvYTr93aQ=",
        version = "v0.68.1",
    )
    go_repository(
        name = "tools_gotest_v3",
        importpath = "gotest.tools/v3",
        sum = "h1:EENdUnS3pdur5nybKYIh2Vfgc8IUNBjxDPSjtiJcOzU=",
        version = "v3.5.1",
    )
